/*
Copyright The Flagger Authors.

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

package v1alpha1

import (
	v1alpha1 "github.com/weaveworks/flagger/pkg/apis/appmesh/v1alpha1"
	scheme "github.com/weaveworks/flagger/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VirtualNodesGetter has a method to return a VirtualNodeInterface.
// A group's client should implement this interface.
type VirtualNodesGetter interface {
	VirtualNodes(namespace string) VirtualNodeInterface
}

// VirtualNodeInterface has methods to work with VirtualNode resources.
type VirtualNodeInterface interface {
	Create(*v1alpha1.VirtualNode) (*v1alpha1.VirtualNode, error)
	Update(*v1alpha1.VirtualNode) (*v1alpha1.VirtualNode, error)
	UpdateStatus(*v1alpha1.VirtualNode) (*v1alpha1.VirtualNode, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VirtualNode, error)
	List(opts v1.ListOptions) (*v1alpha1.VirtualNodeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualNode, err error)
	VirtualNodeExpansion
}

// virtualNodes implements VirtualNodeInterface
type virtualNodes struct {
	client rest.Interface
	ns     string
}

// newVirtualNodes returns a VirtualNodes
func newVirtualNodes(c *AppmeshV1alpha1Client, namespace string) *virtualNodes {
	return &virtualNodes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualNode, and returns the corresponding virtualNode object, and an error if there is any.
func (c *virtualNodes) Get(name string, options v1.GetOptions) (result *v1alpha1.VirtualNode, err error) {
	result = &v1alpha1.VirtualNode{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualnodes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualNodes that match those selectors.
func (c *virtualNodes) List(opts v1.ListOptions) (result *v1alpha1.VirtualNodeList, err error) {
	result = &v1alpha1.VirtualNodeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualnodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualNodes.
func (c *virtualNodes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualnodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a virtualNode and creates it.  Returns the server's representation of the virtualNode, and an error, if there is any.
func (c *virtualNodes) Create(virtualNode *v1alpha1.VirtualNode) (result *v1alpha1.VirtualNode, err error) {
	result = &v1alpha1.VirtualNode{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualnodes").
		Body(virtualNode).
		Do().
		Into(result)
	return
}

// Update takes the representation of a virtualNode and updates it. Returns the server's representation of the virtualNode, and an error, if there is any.
func (c *virtualNodes) Update(virtualNode *v1alpha1.VirtualNode) (result *v1alpha1.VirtualNode, err error) {
	result = &v1alpha1.VirtualNode{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualnodes").
		Name(virtualNode.Name).
		Body(virtualNode).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *virtualNodes) UpdateStatus(virtualNode *v1alpha1.VirtualNode) (result *v1alpha1.VirtualNode, err error) {
	result = &v1alpha1.VirtualNode{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualnodes").
		Name(virtualNode.Name).
		SubResource("status").
		Body(virtualNode).
		Do().
		Into(result)
	return
}

// Delete takes name of the virtualNode and deletes it. Returns an error if one occurs.
func (c *virtualNodes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualnodes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualNodes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualnodes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched virtualNode.
func (c *virtualNodes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualNode, err error) {
	result = &v1alpha1.VirtualNode{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualnodes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
