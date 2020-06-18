/*
Copyright AndreasM009.

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
	"context"

	v1alpha1 "github.com/AndreasM009/eventstore/pkg/apis/eventstore/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEventstores implements EventstoreInterface
type FakeEventstores struct {
	Fake *FakeEventstoreV1alpha1
	ns   string
}

var eventstoresResource = schema.GroupVersionResource{Group: "eventstore.io", Version: "v1alpha1", Resource: "eventstores"}

var eventstoresKind = schema.GroupVersionKind{Group: "eventstore.io", Version: "v1alpha1", Kind: "Eventstore"}

// Get takes name of the eventstore, and returns the corresponding eventstore object, and an error if there is any.
func (c *FakeEventstores) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Eventstore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(eventstoresResource, c.ns, name), &v1alpha1.Eventstore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Eventstore), err
}

// List takes label and field selectors, and returns the list of Eventstores that match those selectors.
func (c *FakeEventstores) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EventstoreList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(eventstoresResource, eventstoresKind, c.ns, opts), &v1alpha1.EventstoreList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EventstoreList{ListMeta: obj.(*v1alpha1.EventstoreList).ListMeta}
	for _, item := range obj.(*v1alpha1.EventstoreList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested eventstores.
func (c *FakeEventstores) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(eventstoresResource, c.ns, opts))

}

// Create takes the representation of a eventstore and creates it.  Returns the server's representation of the eventstore, and an error, if there is any.
func (c *FakeEventstores) Create(ctx context.Context, eventstore *v1alpha1.Eventstore, opts v1.CreateOptions) (result *v1alpha1.Eventstore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(eventstoresResource, c.ns, eventstore), &v1alpha1.Eventstore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Eventstore), err
}

// Update takes the representation of a eventstore and updates it. Returns the server's representation of the eventstore, and an error, if there is any.
func (c *FakeEventstores) Update(ctx context.Context, eventstore *v1alpha1.Eventstore, opts v1.UpdateOptions) (result *v1alpha1.Eventstore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(eventstoresResource, c.ns, eventstore), &v1alpha1.Eventstore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Eventstore), err
}

// Delete takes name of the eventstore and deletes it. Returns an error if one occurs.
func (c *FakeEventstores) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(eventstoresResource, c.ns, name), &v1alpha1.Eventstore{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEventstores) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(eventstoresResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EventstoreList{})
	return err
}

// Patch applies the patch and returns the patched eventstore.
func (c *FakeEventstores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Eventstore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(eventstoresResource, c.ns, name, pt, data, subresources...), &v1alpha1.Eventstore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Eventstore), err
}
