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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "yunion.io/x/onecloud-operator/pkg/apis/onecloud/v1alpha1"
)

// FakeOnecloudClusters implements OnecloudClusterInterface
type FakeOnecloudClusters struct {
	Fake *FakeOnecloudV1alpha1
	ns   string
}

var onecloudclustersResource = schema.GroupVersionResource{Group: "onecloud.yunion.io", Version: "v1alpha1", Resource: "onecloudclusters"}

var onecloudclustersKind = schema.GroupVersionKind{Group: "onecloud.yunion.io", Version: "v1alpha1", Kind: "OnecloudCluster"}

// Get takes name of the onecloudCluster, and returns the corresponding onecloudCluster object, and an error if there is any.
func (c *FakeOnecloudClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.OnecloudCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(onecloudclustersResource, c.ns, name), &v1alpha1.OnecloudCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OnecloudCluster), err
}

// List takes label and field selectors, and returns the list of OnecloudClusters that match those selectors.
func (c *FakeOnecloudClusters) List(opts v1.ListOptions) (result *v1alpha1.OnecloudClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(onecloudclustersResource, onecloudclustersKind, c.ns, opts), &v1alpha1.OnecloudClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OnecloudClusterList{ListMeta: obj.(*v1alpha1.OnecloudClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.OnecloudClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested onecloudClusters.
func (c *FakeOnecloudClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(onecloudclustersResource, c.ns, opts))

}

// Create takes the representation of a onecloudCluster and creates it.  Returns the server's representation of the onecloudCluster, and an error, if there is any.
func (c *FakeOnecloudClusters) Create(onecloudCluster *v1alpha1.OnecloudCluster) (result *v1alpha1.OnecloudCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(onecloudclustersResource, c.ns, onecloudCluster), &v1alpha1.OnecloudCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OnecloudCluster), err
}

// Update takes the representation of a onecloudCluster and updates it. Returns the server's representation of the onecloudCluster, and an error, if there is any.
func (c *FakeOnecloudClusters) Update(onecloudCluster *v1alpha1.OnecloudCluster) (result *v1alpha1.OnecloudCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(onecloudclustersResource, c.ns, onecloudCluster), &v1alpha1.OnecloudCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OnecloudCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOnecloudClusters) UpdateStatus(onecloudCluster *v1alpha1.OnecloudCluster) (*v1alpha1.OnecloudCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(onecloudclustersResource, "status", c.ns, onecloudCluster), &v1alpha1.OnecloudCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OnecloudCluster), err
}

// Delete takes name of the onecloudCluster and deletes it. Returns an error if one occurs.
func (c *FakeOnecloudClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(onecloudclustersResource, c.ns, name), &v1alpha1.OnecloudCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOnecloudClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(onecloudclustersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.OnecloudClusterList{})
	return err
}

// Patch applies the patch and returns the patched onecloudCluster.
func (c *FakeOnecloudClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OnecloudCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(onecloudclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.OnecloudCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OnecloudCluster), err
}