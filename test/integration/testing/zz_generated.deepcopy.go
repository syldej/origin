// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package testing

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestPluginConfig) DeepCopyInto(out *TestPluginConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestPluginConfig.
func (in *TestPluginConfig) DeepCopy() *TestPluginConfig {
	if in == nil {
		return nil
	}
	out := new(TestPluginConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TestPluginConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
