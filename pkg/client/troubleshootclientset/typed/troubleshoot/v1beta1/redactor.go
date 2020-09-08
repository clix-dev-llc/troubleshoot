/*
Copyright 2019 Replicated, Inc..

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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/replicatedhq/troubleshoot/pkg/apis/troubleshoot/v1beta1"
	scheme "github.com/replicatedhq/troubleshoot/pkg/client/troubleshootclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RedactorsGetter has a method to return a RedactorInterface.
// A group's client should implement this interface.
type RedactorsGetter interface {
	Redactors(namespace string) RedactorInterface
}

// RedactorInterface has methods to work with Redactor resources.
type RedactorInterface interface {
	Create(ctx context.Context, redactor *v1beta1.Redactor, opts v1.CreateOptions) (*v1beta1.Redactor, error)
	Update(ctx context.Context, redactor *v1beta1.Redactor, opts v1.UpdateOptions) (*v1beta1.Redactor, error)
	UpdateStatus(ctx context.Context, redactor *v1beta1.Redactor, opts v1.UpdateOptions) (*v1beta1.Redactor, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.Redactor, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.RedactorList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Redactor, err error)
	RedactorExpansion
}

// redactors implements RedactorInterface
type redactors struct {
	client rest.Interface
	ns     string
}

// newRedactors returns a Redactors
func newRedactors(c *TroubleshootV1beta1Client, namespace string) *redactors {
	return &redactors{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the redactor, and returns the corresponding redactor object, and an error if there is any.
func (c *redactors) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Redactor, err error) {
	result = &v1beta1.Redactor{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("redactors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Redactors that match those selectors.
func (c *redactors) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.RedactorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.RedactorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("redactors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested redactors.
func (c *redactors) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("redactors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a redactor and creates it.  Returns the server's representation of the redactor, and an error, if there is any.
func (c *redactors) Create(ctx context.Context, redactor *v1beta1.Redactor, opts v1.CreateOptions) (result *v1beta1.Redactor, err error) {
	result = &v1beta1.Redactor{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("redactors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(redactor).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a redactor and updates it. Returns the server's representation of the redactor, and an error, if there is any.
func (c *redactors) Update(ctx context.Context, redactor *v1beta1.Redactor, opts v1.UpdateOptions) (result *v1beta1.Redactor, err error) {
	result = &v1beta1.Redactor{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("redactors").
		Name(redactor.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(redactor).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *redactors) UpdateStatus(ctx context.Context, redactor *v1beta1.Redactor, opts v1.UpdateOptions) (result *v1beta1.Redactor, err error) {
	result = &v1beta1.Redactor{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("redactors").
		Name(redactor.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(redactor).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the redactor and deletes it. Returns an error if one occurs.
func (c *redactors) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("redactors").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *redactors) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("redactors").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched redactor.
func (c *redactors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Redactor, err error) {
	result = &v1beta1.Redactor{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("redactors").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
