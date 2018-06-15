// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/openshift/origin/pkg/cmd/openshift-operators/apis/webconsole/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOpenShiftWebConsoleConfigs implements OpenShiftWebConsoleConfigInterface
type FakeOpenShiftWebConsoleConfigs struct {
	Fake *FakeWebconsoleV1alpha1
}

var openshiftwebconsoleconfigsResource = schema.GroupVersionResource{Group: "webconsole", Version: "v1alpha1", Resource: "openshiftwebconsoleconfigs"}

var openshiftwebconsoleconfigsKind = schema.GroupVersionKind{Group: "webconsole", Version: "v1alpha1", Kind: "OpenShiftWebConsoleConfig"}

// Get takes name of the openShiftWebConsoleConfig, and returns the corresponding openShiftWebConsoleConfig object, and an error if there is any.
func (c *FakeOpenShiftWebConsoleConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.OpenShiftWebConsoleConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(openshiftwebconsoleconfigsResource, name), &v1alpha1.OpenShiftWebConsoleConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpenShiftWebConsoleConfig), err
}

// List takes label and field selectors, and returns the list of OpenShiftWebConsoleConfigs that match those selectors.
func (c *FakeOpenShiftWebConsoleConfigs) List(opts v1.ListOptions) (result *v1alpha1.OpenShiftWebConsoleConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(openshiftwebconsoleconfigsResource, openshiftwebconsoleconfigsKind, opts), &v1alpha1.OpenShiftWebConsoleConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OpenShiftWebConsoleConfigList{}
	for _, item := range obj.(*v1alpha1.OpenShiftWebConsoleConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested openShiftWebConsoleConfigs.
func (c *FakeOpenShiftWebConsoleConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(openshiftwebconsoleconfigsResource, opts))
}

// Create takes the representation of a openShiftWebConsoleConfig and creates it.  Returns the server's representation of the openShiftWebConsoleConfig, and an error, if there is any.
func (c *FakeOpenShiftWebConsoleConfigs) Create(openShiftWebConsoleConfig *v1alpha1.OpenShiftWebConsoleConfig) (result *v1alpha1.OpenShiftWebConsoleConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(openshiftwebconsoleconfigsResource, openShiftWebConsoleConfig), &v1alpha1.OpenShiftWebConsoleConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpenShiftWebConsoleConfig), err
}

// Update takes the representation of a openShiftWebConsoleConfig and updates it. Returns the server's representation of the openShiftWebConsoleConfig, and an error, if there is any.
func (c *FakeOpenShiftWebConsoleConfigs) Update(openShiftWebConsoleConfig *v1alpha1.OpenShiftWebConsoleConfig) (result *v1alpha1.OpenShiftWebConsoleConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(openshiftwebconsoleconfigsResource, openShiftWebConsoleConfig), &v1alpha1.OpenShiftWebConsoleConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpenShiftWebConsoleConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOpenShiftWebConsoleConfigs) UpdateStatus(openShiftWebConsoleConfig *v1alpha1.OpenShiftWebConsoleConfig) (*v1alpha1.OpenShiftWebConsoleConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(openshiftwebconsoleconfigsResource, "status", openShiftWebConsoleConfig), &v1alpha1.OpenShiftWebConsoleConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpenShiftWebConsoleConfig), err
}

// Delete takes name of the openShiftWebConsoleConfig and deletes it. Returns an error if one occurs.
func (c *FakeOpenShiftWebConsoleConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(openshiftwebconsoleconfigsResource, name), &v1alpha1.OpenShiftWebConsoleConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOpenShiftWebConsoleConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(openshiftwebconsoleconfigsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.OpenShiftWebConsoleConfigList{})
	return err
}

// Patch applies the patch and returns the patched openShiftWebConsoleConfig.
func (c *FakeOpenShiftWebConsoleConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OpenShiftWebConsoleConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(openshiftwebconsoleconfigsResource, name, data, subresources...), &v1alpha1.OpenShiftWebConsoleConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpenShiftWebConsoleConfig), err
}
