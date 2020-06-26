/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/vega-project/ccb-operator/pkg/apis/calculations/v1"
	scheme "github.com/vega-project/ccb-operator/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CalculationsGetter has a method to return a CalculationInterface.
// A group's client should implement this interface.
type CalculationsGetter interface {
	Calculations() CalculationInterface
}

// CalculationInterface has methods to work with Calculation resources.
type CalculationInterface interface {
	Create(ctx context.Context, calculation *v1.Calculation, opts metav1.CreateOptions) (*v1.Calculation, error)
	Update(ctx context.Context, calculation *v1.Calculation, opts metav1.UpdateOptions) (*v1.Calculation, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Calculation, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.CalculationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Calculation, err error)
	CalculationExpansion
}

// calculations implements CalculationInterface
type calculations struct {
	client rest.Interface
}

// newCalculations returns a Calculations
func newCalculations(c *VegaV1Client) *calculations {
	return &calculations{
		client: c.RESTClient(),
	}
}

// Get takes name of the calculation, and returns the corresponding calculation object, and an error if there is any.
func (c *calculations) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Calculation, err error) {
	result = &v1.Calculation{}
	err = c.client.Get().
		Resource("calculations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Calculations that match those selectors.
func (c *calculations) List(ctx context.Context, opts metav1.ListOptions) (result *v1.CalculationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.CalculationList{}
	err = c.client.Get().
		Resource("calculations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested calculations.
func (c *calculations) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("calculations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a calculation and creates it.  Returns the server's representation of the calculation, and an error, if there is any.
func (c *calculations) Create(ctx context.Context, calculation *v1.Calculation, opts metav1.CreateOptions) (result *v1.Calculation, err error) {
	result = &v1.Calculation{}
	err = c.client.Post().
		Resource("calculations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(calculation).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a calculation and updates it. Returns the server's representation of the calculation, and an error, if there is any.
func (c *calculations) Update(ctx context.Context, calculation *v1.Calculation, opts metav1.UpdateOptions) (result *v1.Calculation, err error) {
	result = &v1.Calculation{}
	err = c.client.Put().
		Resource("calculations").
		Name(calculation.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(calculation).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the calculation and deletes it. Returns an error if one occurs.
func (c *calculations) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("calculations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *calculations) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("calculations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched calculation.
func (c *calculations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Calculation, err error) {
	result = &v1.Calculation{}
	err = c.client.Patch(pt).
		Resource("calculations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
