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

package v1alpha1

import (
	"github.com/fluxcd/pkg/apis/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Decryption) DeepCopyInto(out *Decryption) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(meta.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Decryption.
func (in *Decryption) DeepCopy() *Decryption {
	if in == nil {
		return nil
	}
	out := new(Decryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DurationOrNever) DeepCopyInto(out *DurationOrNever) {
	*out = *in
	out.Duration = in.Duration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DurationOrNever.
func (in *DurationOrNever) DeepCopy() *DurationOrNever {
	if in == nil {
		return nil
	}
	out := new(DurationOrNever)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FixedImage) DeepCopyInto(out *FixedImage) {
	*out = *in
	if in.DeployedImage != nil {
		in, out := &in.DeployedImage, &out.DeployedImage
		*out = new(string)
		**out = **in
	}
	if in.RegistryImage != nil {
		in, out := &in.RegistryImage, &out.RegistryImage
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Object != nil {
		in, out := &in.Object, &out.Object
		*out = new(ObjectRef)
		**out = **in
	}
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = new(string)
		**out = **in
	}
	if in.Container != nil {
		in, out := &in.Container, &out.Container
		*out = new(string)
		**out = **in
	}
	if in.VersionFilter != nil {
		in, out := &in.VersionFilter, &out.VersionFilter
		*out = new(string)
		**out = **in
	}
	if in.DeployTags != nil {
		in, out := &in.DeployTags, &out.DeployTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DeploymentDir != nil {
		in, out := &in.DeploymentDir, &out.DeploymentDir
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FixedImage.
func (in *FixedImage) DeepCopy() *FixedImage {
	if in == nil {
		return nil
	}
	out := new(FixedImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KluctlDeployment) DeepCopyInto(out *KluctlDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KluctlDeployment.
func (in *KluctlDeployment) DeepCopy() *KluctlDeployment {
	if in == nil {
		return nil
	}
	out := new(KluctlDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KluctlDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KluctlDeploymentList) DeepCopyInto(out *KluctlDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KluctlDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KluctlDeploymentList.
func (in *KluctlDeploymentList) DeepCopy() *KluctlDeploymentList {
	if in == nil {
		return nil
	}
	out := new(KluctlDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KluctlDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KluctlDeploymentSpec) DeepCopyInto(out *KluctlDeploymentSpec) {
	*out = *in
	out.SourceRef = in.SourceRef
	if in.Decryption != nil {
		in, out := &in.Decryption, &out.Decryption
		*out = new(Decryption)
		(*in).DeepCopyInto(*out)
	}
	out.Interval = in.Interval
	if in.RetryInterval != nil {
		in, out := &in.RetryInterval, &out.RetryInterval
		*out = new(v1.Duration)
		**out = **in
	}
	if in.DeployInterval != nil {
		in, out := &in.DeployInterval, &out.DeployInterval
		*out = new(DurationOrNever)
		**out = **in
	}
	if in.ValidateInterval != nil {
		in, out := &in.ValidateInterval, &out.ValidateInterval
		*out = new(DurationOrNever)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.RegistrySecrets != nil {
		in, out := &in.RegistrySecrets, &out.RegistrySecrets
		*out = make([]meta.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.HelmCredentials != nil {
		in, out := &in.HelmCredentials, &out.HelmCredentials
		*out = make([]meta.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.KubeConfig != nil {
		in, out := &in.KubeConfig, &out.KubeConfig
		*out = new(KubeConfig)
		**out = **in
	}
	if in.RenameContexts != nil {
		in, out := &in.RenameContexts, &out.RenameContexts
		*out = make([]RenameContext, len(*in))
		copy(*out, *in)
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
	if in.TargetNameOverride != nil {
		in, out := &in.TargetNameOverride, &out.TargetNameOverride
		*out = new(string)
		**out = **in
	}
	if in.Context != nil {
		in, out := &in.Context, &out.Context
		*out = new(string)
		**out = **in
	}
	in.Args.DeepCopyInto(&out.Args)
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]FixedImage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IncludeTags != nil {
		in, out := &in.IncludeTags, &out.IncludeTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeTags != nil {
		in, out := &in.ExcludeTags, &out.ExcludeTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IncludeDeploymentDirs != nil {
		in, out := &in.IncludeDeploymentDirs, &out.IncludeDeploymentDirs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeDeploymentDirs != nil {
		in, out := &in.ExcludeDeploymentDirs, &out.ExcludeDeploymentDirs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KluctlDeploymentSpec.
func (in *KluctlDeploymentSpec) DeepCopy() *KluctlDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(KluctlDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KluctlDeploymentStatus) DeepCopyInto(out *KluctlDeploymentStatus) {
	*out = *in
	out.ReconcileRequestStatus = in.ReconcileRequestStatus
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastDeployResult != nil {
		in, out := &in.LastDeployResult, &out.LastDeployResult
		*out = new(LastCommandResult)
		(*in).DeepCopyInto(*out)
	}
	if in.LastPruneResult != nil {
		in, out := &in.LastPruneResult, &out.LastPruneResult
		*out = new(LastCommandResult)
		(*in).DeepCopyInto(*out)
	}
	if in.LastValidateResult != nil {
		in, out := &in.LastValidateResult, &out.LastValidateResult
		*out = new(LastValidateResult)
		(*in).DeepCopyInto(*out)
	}
	if in.CommonLabels != nil {
		in, out := &in.CommonLabels, &out.CommonLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RawTarget != nil {
		in, out := &in.RawTarget, &out.RawTarget
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KluctlDeploymentStatus.
func (in *KluctlDeploymentStatus) DeepCopy() *KluctlDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(KluctlDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeConfig) DeepCopyInto(out *KubeConfig) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeConfig.
func (in *KubeConfig) DeepCopy() *KubeConfig {
	if in == nil {
		return nil
	}
	out := new(KubeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LastCommandResult) DeepCopyInto(out *LastCommandResult) {
	*out = *in
	in.ReconcileResultBase.DeepCopyInto(&out.ReconcileResultBase)
	if in.RawResult != nil {
		in, out := &in.RawResult, &out.RawResult
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LastCommandResult.
func (in *LastCommandResult) DeepCopy() *LastCommandResult {
	if in == nil {
		return nil
	}
	out := new(LastCommandResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LastValidateResult) DeepCopyInto(out *LastValidateResult) {
	*out = *in
	in.ReconcileResultBase.DeepCopyInto(&out.ReconcileResultBase)
	if in.RawResult != nil {
		in, out := &in.RawResult, &out.RawResult
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LastValidateResult.
func (in *LastValidateResult) DeepCopy() *LastValidateResult {
	if in == nil {
		return nil
	}
	out := new(LastValidateResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectRef) DeepCopyInto(out *ObjectRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectRef.
func (in *ObjectRef) DeepCopy() *ObjectRef {
	if in == nil {
		return nil
	}
	out := new(ObjectRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReconcileResultBase) DeepCopyInto(out *ReconcileResultBase) {
	*out = *in
	in.AttemptedAt.DeepCopyInto(&out.AttemptedAt)
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
	if in.TargetNameOverride != nil {
		in, out := &in.TargetNameOverride, &out.TargetNameOverride
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReconcileResultBase.
func (in *ReconcileResultBase) DeepCopy() *ReconcileResultBase {
	if in == nil {
		return nil
	}
	out := new(ReconcileResultBase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RenameContext) DeepCopyInto(out *RenameContext) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RenameContext.
func (in *RenameContext) DeepCopy() *RenameContext {
	if in == nil {
		return nil
	}
	out := new(RenameContext)
	in.DeepCopyInto(out)
	return out
}
