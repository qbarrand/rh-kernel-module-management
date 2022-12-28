//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Build) DeepCopyInto(out *Build) {
	*out = *in
	if in.BuildArgs != nil {
		in, out := &in.BuildArgs, &out.BuildArgs
		*out = make([]BuildArg, len(*in))
		copy(*out, *in)
	}
	if in.DockerfileConfigMap != nil {
		in, out := &in.DockerfileConfigMap, &out.DockerfileConfigMap
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	out.BaseImageRegistryTLS = in.BaseImageRegistryTLS
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.KanikoParams != nil {
		in, out := &in.KanikoParams, &out.KanikoParams
		*out = new(KanikoParams)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Build.
func (in *Build) DeepCopy() *Build {
	if in == nil {
		return nil
	}
	out := new(Build)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildArg) DeepCopyInto(out *BuildArg) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildArg.
func (in *BuildArg) DeepCopy() *BuildArg {
	if in == nil {
		return nil
	}
	out := new(BuildArg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CRStatus) DeepCopyInto(out *CRStatus) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CRStatus.
func (in *CRStatus) DeepCopy() *CRStatus {
	if in == nil {
		return nil
	}
	out := new(CRStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonSetStatus) DeepCopyInto(out *DaemonSetStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonSetStatus.
func (in *DaemonSetStatus) DeepCopy() *DaemonSetStatus {
	if in == nil {
		return nil
	}
	out := new(DaemonSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePluginContainerSpec) DeepCopyInto(out *DevicePluginContainerSpec) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePluginContainerSpec.
func (in *DevicePluginContainerSpec) DeepCopy() *DevicePluginContainerSpec {
	if in == nil {
		return nil
	}
	out := new(DevicePluginContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePluginSpec) DeepCopyInto(out *DevicePluginSpec) {
	*out = *in
	in.Container.DeepCopyInto(&out.Container)
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePluginSpec.
func (in *DevicePluginSpec) DeepCopy() *DevicePluginSpec {
	if in == nil {
		return nil
	}
	out := new(DevicePluginSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KanikoParams) DeepCopyInto(out *KanikoParams) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KanikoParams.
func (in *KanikoParams) DeepCopy() *KanikoParams {
	if in == nil {
		return nil
	}
	out := new(KanikoParams)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KernelMapping) DeepCopyInto(out *KernelMapping) {
	*out = *in
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(Build)
		(*in).DeepCopyInto(*out)
	}
	if in.Sign != nil {
		in, out := &in.Sign, &out.Sign
		*out = new(Sign)
		(*in).DeepCopyInto(*out)
	}
	if in.RegistryTLS != nil {
		in, out := &in.RegistryTLS, &out.RegistryTLS
		*out = new(TLSOptions)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KernelMapping.
func (in *KernelMapping) DeepCopy() *KernelMapping {
	if in == nil {
		return nil
	}
	out := new(KernelMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModprobeArgs) DeepCopyInto(out *ModprobeArgs) {
	*out = *in
	if in.Load != nil {
		in, out := &in.Load, &out.Load
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Unload != nil {
		in, out := &in.Unload, &out.Unload
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModprobeArgs.
func (in *ModprobeArgs) DeepCopy() *ModprobeArgs {
	if in == nil {
		return nil
	}
	out := new(ModprobeArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModprobeSpec) DeepCopyInto(out *ModprobeSpec) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = new(ModprobeArgs)
		(*in).DeepCopyInto(*out)
	}
	if in.RawArgs != nil {
		in, out := &in.RawArgs, &out.RawArgs
		*out = new(ModprobeArgs)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModprobeSpec.
func (in *ModprobeSpec) DeepCopy() *ModprobeSpec {
	if in == nil {
		return nil
	}
	out := new(ModprobeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Module) DeepCopyInto(out *Module) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Module.
func (in *Module) DeepCopy() *Module {
	if in == nil {
		return nil
	}
	out := new(Module)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Module) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleList) DeepCopyInto(out *ModuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Module, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleList.
func (in *ModuleList) DeepCopy() *ModuleList {
	if in == nil {
		return nil
	}
	out := new(ModuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleLoaderContainerSpec) DeepCopyInto(out *ModuleLoaderContainerSpec) {
	*out = *in
	if in.Build != nil {
		in, out := &in.Build, &out.Build
		*out = new(Build)
		(*in).DeepCopyInto(*out)
	}
	if in.Sign != nil {
		in, out := &in.Sign, &out.Sign
		*out = new(Sign)
		(*in).DeepCopyInto(*out)
	}
	if in.KernelMappings != nil {
		in, out := &in.KernelMappings, &out.KernelMappings
		*out = make([]KernelMapping, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Modprobe.DeepCopyInto(&out.Modprobe)
	out.RegistryTLS = in.RegistryTLS
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleLoaderContainerSpec.
func (in *ModuleLoaderContainerSpec) DeepCopy() *ModuleLoaderContainerSpec {
	if in == nil {
		return nil
	}
	out := new(ModuleLoaderContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleLoaderSpec) DeepCopyInto(out *ModuleLoaderSpec) {
	*out = *in
	in.Container.DeepCopyInto(&out.Container)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleLoaderSpec.
func (in *ModuleLoaderSpec) DeepCopy() *ModuleLoaderSpec {
	if in == nil {
		return nil
	}
	out := new(ModuleLoaderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleSpec) DeepCopyInto(out *ModuleSpec) {
	*out = *in
	if in.DevicePlugin != nil {
		in, out := &in.DevicePlugin, &out.DevicePlugin
		*out = new(DevicePluginSpec)
		(*in).DeepCopyInto(*out)
	}
	in.ModuleLoader.DeepCopyInto(&out.ModuleLoader)
	if in.ImageRepoSecret != nil {
		in, out := &in.ImageRepoSecret, &out.ImageRepoSecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleSpec.
func (in *ModuleSpec) DeepCopy() *ModuleSpec {
	if in == nil {
		return nil
	}
	out := new(ModuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleStatus) DeepCopyInto(out *ModuleStatus) {
	*out = *in
	out.DevicePlugin = in.DevicePlugin
	out.ModuleLoader = in.ModuleLoader
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleStatus.
func (in *ModuleStatus) DeepCopy() *ModuleStatus {
	if in == nil {
		return nil
	}
	out := new(ModuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreflightValidation) DeepCopyInto(out *PreflightValidation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreflightValidation.
func (in *PreflightValidation) DeepCopy() *PreflightValidation {
	if in == nil {
		return nil
	}
	out := new(PreflightValidation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PreflightValidation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreflightValidationList) DeepCopyInto(out *PreflightValidationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PreflightValidation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreflightValidationList.
func (in *PreflightValidationList) DeepCopy() *PreflightValidationList {
	if in == nil {
		return nil
	}
	out := new(PreflightValidationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PreflightValidationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreflightValidationOCP) DeepCopyInto(out *PreflightValidationOCP) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreflightValidationOCP.
func (in *PreflightValidationOCP) DeepCopy() *PreflightValidationOCP {
	if in == nil {
		return nil
	}
	out := new(PreflightValidationOCP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PreflightValidationOCP) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreflightValidationOCPList) DeepCopyInto(out *PreflightValidationOCPList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PreflightValidationOCP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreflightValidationOCPList.
func (in *PreflightValidationOCPList) DeepCopy() *PreflightValidationOCPList {
	if in == nil {
		return nil
	}
	out := new(PreflightValidationOCPList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PreflightValidationOCPList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreflightValidationOCPSpec) DeepCopyInto(out *PreflightValidationOCPSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreflightValidationOCPSpec.
func (in *PreflightValidationOCPSpec) DeepCopy() *PreflightValidationOCPSpec {
	if in == nil {
		return nil
	}
	out := new(PreflightValidationOCPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreflightValidationSpec) DeepCopyInto(out *PreflightValidationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreflightValidationSpec.
func (in *PreflightValidationSpec) DeepCopy() *PreflightValidationSpec {
	if in == nil {
		return nil
	}
	out := new(PreflightValidationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreflightValidationStatus) DeepCopyInto(out *PreflightValidationStatus) {
	*out = *in
	if in.CRStatuses != nil {
		in, out := &in.CRStatuses, &out.CRStatuses
		*out = make(map[string]*CRStatus, len(*in))
		for key, val := range *in {
			var outVal *CRStatus
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(CRStatus)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreflightValidationStatus.
func (in *PreflightValidationStatus) DeepCopy() *PreflightValidationStatus {
	if in == nil {
		return nil
	}
	out := new(PreflightValidationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sign) DeepCopyInto(out *Sign) {
	*out = *in
	out.UnsignedImageRegistryTLS = in.UnsignedImageRegistryTLS
	if in.KeySecret != nil {
		in, out := &in.KeySecret, &out.KeySecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.CertSecret != nil {
		in, out := &in.CertSecret, &out.CertSecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.FilesToSign != nil {
		in, out := &in.FilesToSign, &out.FilesToSign
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sign.
func (in *Sign) DeepCopy() *Sign {
	if in == nil {
		return nil
	}
	out := new(Sign)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSOptions) DeepCopyInto(out *TLSOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSOptions.
func (in *TLSOptions) DeepCopy() *TLSOptions {
	if in == nil {
		return nil
	}
	out := new(TLSOptions)
	in.DeepCopyInto(out)
	return out
}
