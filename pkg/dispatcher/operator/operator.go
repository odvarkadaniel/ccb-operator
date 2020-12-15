package operator

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/runtime"

	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"

	clientset "github.com/vega-project/ccb-operator/pkg/client/clientset/versioned"
	calculationscheme "github.com/vega-project/ccb-operator/pkg/client/clientset/versioned/scheme"
	informers "github.com/vega-project/ccb-operator/pkg/client/informers/externalversions"
	"github.com/vega-project/ccb-operator/pkg/dispatcher/operator/calculations"
	"github.com/vega-project/ccb-operator/pkg/dispatcher/operator/workers"
)

type Operator struct {
	ctx                    context.Context
	logger                 *logrus.Logger
	kubeclientset          kubernetes.Interface
	vegaclientset          clientset.Interface
	kubeInformer           kubeinformers.SharedInformerFactory
	informer               informers.SharedInformerFactory
	calculationsController *calculations.Controller
	podsController         *workers.Controller
	redisURL               string
	redisPW                string
}

// NewMainOperator return a new Operator
func NewMainOperator(ctx context.Context, kubeclientset kubernetes.Interface, vegaclientset clientset.Interface, redisURL string, redisPW string) *Operator {
	logger := logrus.New()
	logger.Level = logrus.DebugLevel
	if dataPW, err := ioutil.ReadFile("redis.conf"); err != nil { //wasnt able to find the path to the file/not sure where it is
		fmt.Errorf("Failed to retrieve database password from a file: %s", err.Error())
		return &Operator{} //not really sure about that
	}
	return &Operator{
		ctx:           ctx,
		logger:        logger,
		kubeclientset: kubeclientset,
		vegaclientset: vegaclientset,
		redisURL:      redisURL,
		redisPW:       dataPW,
	}
}

// Initialize initializes the operator with both calculation/pods controllers and informers.
func (op *Operator) Initialize() {
	op.kubeInformer = kubeinformers.NewSharedInformerFactoryWithOptions(op.kubeclientset, 30*time.Second,
		kubeinformers.WithTweakListOptions(func(options *metav1.ListOptions) {
			options.LabelSelector = fields.Set{"name": "vega-worker"}.AsSelector().String()
		}), kubeinformers.WithNamespace("vega"))

	op.informer = informers.NewSharedInformerFactoryWithOptions(op.vegaclientset, 30*time.Second, informers.WithNamespace("vega"))
	runtime.Must(calculationscheme.AddToScheme(scheme.Scheme))

	redisClient := redis.NewClient(&redis.Options{
		Addr:     op.redisURL,
		Password: op.redisPW,
		DB:       0,
	})

	op.calculationsController = calculations.NewController(op.vegaclientset, op.informer.Vega().V1().Calculations(), redisClient)
	op.podsController = workers.NewController(op.ctx, op.kubeclientset, op.kubeInformer.Core().V1().Pods(), op.vegaclientset, op.informer.Vega().V1().Calculations().Lister(), redisClient)
}

// Run starts the calculation and pod controllers.
func (op *Operator) Run(stopCh <-chan struct{}) error {
	op.kubeInformer.Start(stopCh)
	op.informer.Start(stopCh)

	var err error
	go func() { err = op.calculationsController.Run(stopCh) }()
	if err != nil {
		return fmt.Errorf("failed to run Calculations controller: %s", err.Error())
	}

	go func() { err = op.podsController.Run(stopCh) }()
	if err != nil {
		return fmt.Errorf("failed to run Pod controller: %s", err.Error())
	}
	<-stopCh
	op.logger.Info("Shutting down controllers")
	return nil
}
