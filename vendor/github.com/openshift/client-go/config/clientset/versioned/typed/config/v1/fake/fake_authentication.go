// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/openshift/api/config/v1"
	configv1 "github.com/openshift/client-go/config/applyconfigurations/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAuthentications implements AuthenticationInterface
type FakeAuthentications struct {
	Fake *FakeConfigV1
}

var authenticationsResource = v1.SchemeGroupVersion.WithResource("authentications")

var authenticationsKind = v1.SchemeGroupVersion.WithKind("Authentication")

// Get takes name of the authentication, and returns the corresponding authentication object, and an error if there is any.
func (c *FakeAuthentications) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Authentication, err error) {
	emptyResult := &v1.Authentication{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(authenticationsResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Authentication), err
}

// List takes label and field selectors, and returns the list of Authentications that match those selectors.
func (c *FakeAuthentications) List(ctx context.Context, opts metav1.ListOptions) (result *v1.AuthenticationList, err error) {
	emptyResult := &v1.AuthenticationList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(authenticationsResource, authenticationsKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.AuthenticationList{ListMeta: obj.(*v1.AuthenticationList).ListMeta}
	for _, item := range obj.(*v1.AuthenticationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested authentications.
func (c *FakeAuthentications) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(authenticationsResource, opts))
}

// Create takes the representation of a authentication and creates it.  Returns the server's representation of the authentication, and an error, if there is any.
func (c *FakeAuthentications) Create(ctx context.Context, authentication *v1.Authentication, opts metav1.CreateOptions) (result *v1.Authentication, err error) {
	emptyResult := &v1.Authentication{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(authenticationsResource, authentication, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Authentication), err
}

// Update takes the representation of a authentication and updates it. Returns the server's representation of the authentication, and an error, if there is any.
func (c *FakeAuthentications) Update(ctx context.Context, authentication *v1.Authentication, opts metav1.UpdateOptions) (result *v1.Authentication, err error) {
	emptyResult := &v1.Authentication{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(authenticationsResource, authentication, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Authentication), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAuthentications) UpdateStatus(ctx context.Context, authentication *v1.Authentication, opts metav1.UpdateOptions) (result *v1.Authentication, err error) {
	emptyResult := &v1.Authentication{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(authenticationsResource, "status", authentication, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Authentication), err
}

// Delete takes name of the authentication and deletes it. Returns an error if one occurs.
func (c *FakeAuthentications) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(authenticationsResource, name, opts), &v1.Authentication{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAuthentications) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(authenticationsResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.AuthenticationList{})
	return err
}

// Patch applies the patch and returns the patched authentication.
func (c *FakeAuthentications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Authentication, err error) {
	emptyResult := &v1.Authentication{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(authenticationsResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Authentication), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied authentication.
func (c *FakeAuthentications) Apply(ctx context.Context, authentication *configv1.AuthenticationApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Authentication, err error) {
	if authentication == nil {
		return nil, fmt.Errorf("authentication provided to Apply must not be nil")
	}
	data, err := json.Marshal(authentication)
	if err != nil {
		return nil, err
	}
	name := authentication.Name
	if name == nil {
		return nil, fmt.Errorf("authentication.Name must be provided to Apply")
	}
	emptyResult := &v1.Authentication{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(authenticationsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Authentication), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeAuthentications) ApplyStatus(ctx context.Context, authentication *configv1.AuthenticationApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Authentication, err error) {
	if authentication == nil {
		return nil, fmt.Errorf("authentication provided to Apply must not be nil")
	}
	data, err := json.Marshal(authentication)
	if err != nil {
		return nil, err
	}
	name := authentication.Name
	if name == nil {
		return nil, fmt.Errorf("authentication.Name must be provided to Apply")
	}
	emptyResult := &v1.Authentication{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(authenticationsResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.Authentication), err
}
