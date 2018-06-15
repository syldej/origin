// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "github.com/openshift/origin/pkg/cmd/openshift-operators/generated/clientset/versioned"
	webconsolev1alpha1 "github.com/openshift/origin/pkg/cmd/openshift-operators/generated/clientset/versioned/typed/webconsole/v1alpha1"
	fakewebconsolev1alpha1 "github.com/openshift/origin/pkg/cmd/openshift-operators/generated/clientset/versioned/typed/webconsole/v1alpha1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	fakePtr := testing.Fake{}
	fakePtr.AddReactor("*", "*", testing.ObjectReaction(o))
	fakePtr.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return &Clientset{fakePtr, &fakediscovery.FakeDiscovery{Fake: &fakePtr}}
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

var _ clientset.Interface = &Clientset{}

// WebconsoleV1alpha1 retrieves the WebconsoleV1alpha1Client
func (c *Clientset) WebconsoleV1alpha1() webconsolev1alpha1.WebconsoleV1alpha1Interface {
	return &fakewebconsolev1alpha1.FakeWebconsoleV1alpha1{Fake: &c.Fake}
}

// Webconsole retrieves the WebconsoleV1alpha1Client
func (c *Clientset) Webconsole() webconsolev1alpha1.WebconsoleV1alpha1Interface {
	return &fakewebconsolev1alpha1.FakeWebconsoleV1alpha1{Fake: &c.Fake}
}
