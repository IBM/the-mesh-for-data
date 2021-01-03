// +build !ignore_autogenerated

// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	protobuf "github.com/ibm/the-mesh-for-data/pkg/connectors/protobuf"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationDetails) DeepCopyInto(out *ApplicationDetails) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationDetails.
func (in *ApplicationDetails) DeepCopy() *ApplicationDetails {
	if in == nil {
		return nil
	}
	out := new(ApplicationDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Blueprint) DeepCopyInto(out *Blueprint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Blueprint.
func (in *Blueprint) DeepCopy() *Blueprint {
	if in == nil {
		return nil
	}
	out := new(Blueprint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Blueprint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlueprintList) DeepCopyInto(out *BlueprintList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Blueprint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlueprintList.
func (in *BlueprintList) DeepCopy() *BlueprintList {
	if in == nil {
		return nil
	}
	out := new(BlueprintList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BlueprintList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlueprintSpec) DeepCopyInto(out *BlueprintSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	in.Flow.DeepCopyInto(&out.Flow)
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make([]ComponentTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlueprintSpec.
func (in *BlueprintSpec) DeepCopy() *BlueprintSpec {
	if in == nil {
		return nil
	}
	out := new(BlueprintSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlueprintStatus) DeepCopyInto(out *BlueprintStatus) {
	*out = *in
	out.ObservedState = in.ObservedState
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlueprintStatus.
func (in *BlueprintStatus) DeepCopy() *BlueprintStatus {
	if in == nil {
		return nil
	}
	out := new(BlueprintStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Capability) DeepCopyInto(out *Capability) {
	*out = *in
	if in.SupportedInterfaces != nil {
		in, out := &in.SupportedInterfaces, &out.SupportedInterfaces
		*out = make([]ModuleInOut, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.API != nil {
		in, out := &in.API, &out.API
		*out = new(InterfaceDetails)
		**out = **in
	}
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]protobuf.EnforcementAction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Capability.
func (in *Capability) DeepCopy() *Capability {
	if in == nil {
		return nil
	}
	out := new(Capability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartSpec) DeepCopyInto(out *ChartSpec) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartSpec.
func (in *ChartSpec) DeepCopy() *ChartSpec {
	if in == nil {
		return nil
	}
	out := new(ChartSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentTemplate) DeepCopyInto(out *ComponentTemplate) {
	*out = *in
	in.Chart.DeepCopyInto(&out.Chart)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentTemplate.
func (in *ComponentTemplate) DeepCopy() *ComponentTemplate {
	if in == nil {
		return nil
	}
	out := new(ComponentTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CopyModuleArgs) DeepCopyInto(out *CopyModuleArgs) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	in.Destination.DeepCopyInto(&out.Destination)
	if in.Transformations != nil {
		in, out := &in.Transformations, &out.Transformations
		*out = make([]protobuf.EnforcementAction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CopyModuleArgs.
func (in *CopyModuleArgs) DeepCopy() *CopyModuleArgs {
	if in == nil {
		return nil
	}
	out := new(CopyModuleArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataContext) DeepCopyInto(out *DataContext) {
	*out = *in
	out.IFdetails = in.IFdetails
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataContext.
func (in *DataContext) DeepCopy() *DataContext {
	if in == nil {
		return nil
	}
	out := new(DataContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataFlow) DeepCopyInto(out *DataFlow) {
	*out = *in
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]FlowStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataFlow.
func (in *DataFlow) DeepCopy() *DataFlow {
	if in == nil {
		return nil
	}
	out := new(DataFlow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataStore) DeepCopyInto(out *DataStore) {
	*out = *in
	if in.Connection != nil {
		in, out := &in.Connection, &out.Connection
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataStore.
func (in *DataStore) DeepCopy() *DataStore {
	if in == nil {
		return nil
	}
	out := new(DataStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dependency) DeepCopyInto(out *Dependency) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dependency.
func (in *Dependency) DeepCopy() *Dependency {
	if in == nil {
		return nil
	}
	out := new(Dependency)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowStep) DeepCopyInto(out *FlowStep) {
	*out = *in
	in.Arguments.DeepCopyInto(&out.Arguments)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowStep.
func (in *FlowStep) DeepCopy() *FlowStep {
	if in == nil {
		return nil
	}
	out := new(FlowStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceDetails) DeepCopyInto(out *InterfaceDetails) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceDetails.
func (in *InterfaceDetails) DeepCopy() *InterfaceDetails {
	if in == nil {
		return nil
	}
	out := new(InterfaceDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M4DApplication) DeepCopyInto(out *M4DApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M4DApplication.
func (in *M4DApplication) DeepCopy() *M4DApplication {
	if in == nil {
		return nil
	}
	out := new(M4DApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *M4DApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M4DApplicationList) DeepCopyInto(out *M4DApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]M4DApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M4DApplicationList.
func (in *M4DApplicationList) DeepCopy() *M4DApplicationList {
	if in == nil {
		return nil
	}
	out := new(M4DApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *M4DApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M4DApplicationSpec) DeepCopyInto(out *M4DApplicationSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	out.AppInfo = in.AppInfo
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]DataContext, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M4DApplicationSpec.
func (in *M4DApplicationSpec) DeepCopy() *M4DApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(M4DApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M4DApplicationStatus) DeepCopyInto(out *M4DApplicationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		copy(*out, *in)
	}
	if in.Generated != nil {
		in, out := &in.Generated, &out.Generated
		*out = new(ResourceReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M4DApplicationStatus.
func (in *M4DApplicationStatus) DeepCopy() *M4DApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(M4DApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M4DBucket) DeepCopyInto(out *M4DBucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M4DBucket.
func (in *M4DBucket) DeepCopy() *M4DBucket {
	if in == nil {
		return nil
	}
	out := new(M4DBucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *M4DBucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M4DBucketList) DeepCopyInto(out *M4DBucketList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]M4DBucket, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M4DBucketList.
func (in *M4DBucketList) DeepCopy() *M4DBucketList {
	if in == nil {
		return nil
	}
	out := new(M4DBucketList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *M4DBucketList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M4DBucketSpec) DeepCopyInto(out *M4DBucketSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M4DBucketSpec.
func (in *M4DBucketSpec) DeepCopy() *M4DBucketSpec {
	if in == nil {
		return nil
	}
	out := new(M4DBucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M4DBucketStatus) DeepCopyInto(out *M4DBucketStatus) {
	*out = *in
	if in.Owners != nil {
		in, out := &in.Owners, &out.Owners
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AssetPrefixPerDataset != nil {
		in, out := &in.AssetPrefixPerDataset, &out.AssetPrefixPerDataset
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M4DBucketStatus.
func (in *M4DBucketStatus) DeepCopy() *M4DBucketStatus {
	if in == nil {
		return nil
	}
	out := new(M4DBucketStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M4DModule) DeepCopyInto(out *M4DModule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M4DModule.
func (in *M4DModule) DeepCopy() *M4DModule {
	if in == nil {
		return nil
	}
	out := new(M4DModule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *M4DModule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M4DModuleList) DeepCopyInto(out *M4DModuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]M4DModule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M4DModuleList.
func (in *M4DModuleList) DeepCopy() *M4DModuleList {
	if in == nil {
		return nil
	}
	out := new(M4DModuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *M4DModuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *M4DModuleSpec) DeepCopyInto(out *M4DModuleSpec) {
	*out = *in
	if in.Flows != nil {
		in, out := &in.Flows, &out.Flows
		*out = make([]ModuleFlow, len(*in))
		copy(*out, *in)
	}
	if in.Dependencies != nil {
		in, out := &in.Dependencies, &out.Dependencies
		*out = make([]Dependency, len(*in))
		copy(*out, *in)
	}
	in.Capabilities.DeepCopyInto(&out.Capabilities)
	in.Chart.DeepCopyInto(&out.Chart)
	if in.StatusIndicators != nil {
		in, out := &in.StatusIndicators, &out.StatusIndicators
		*out = make([]ResourceStatusIndicator, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new M4DModuleSpec.
func (in *M4DModuleSpec) DeepCopy() *M4DModuleSpec {
	if in == nil {
		return nil
	}
	out := new(M4DModuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetaBlueprint) DeepCopyInto(out *MetaBlueprint) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetaBlueprint.
func (in *MetaBlueprint) DeepCopy() *MetaBlueprint {
	if in == nil {
		return nil
	}
	out := new(MetaBlueprint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleArguments) DeepCopyInto(out *ModuleArguments) {
	*out = *in
	if in.Copy != nil {
		in, out := &in.Copy, &out.Copy
		*out = new(CopyModuleArgs)
		(*in).DeepCopyInto(*out)
	}
	if in.Read != nil {
		in, out := &in.Read, &out.Read
		*out = make([]ReadModuleArgs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Write != nil {
		in, out := &in.Write, &out.Write
		*out = make([]WriteModuleArgs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleArguments.
func (in *ModuleArguments) DeepCopy() *ModuleArguments {
	if in == nil {
		return nil
	}
	out := new(ModuleArguments)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleInOut) DeepCopyInto(out *ModuleInOut) {
	*out = *in
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(InterfaceDetails)
		**out = **in
	}
	if in.Sink != nil {
		in, out := &in.Sink, &out.Sink
		*out = new(InterfaceDetails)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleInOut.
func (in *ModuleInOut) DeepCopy() *ModuleInOut {
	if in == nil {
		return nil
	}
	out := new(ModuleInOut)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservedState) DeepCopyInto(out *ObservedState) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservedState.
func (in *ObservedState) DeepCopy() *ObservedState {
	if in == nil {
		return nil
	}
	out := new(ObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plotter) DeepCopyInto(out *Plotter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plotter.
func (in *Plotter) DeepCopy() *Plotter {
	if in == nil {
		return nil
	}
	out := new(Plotter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Plotter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlotterList) DeepCopyInto(out *PlotterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Plotter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlotterList.
func (in *PlotterList) DeepCopy() *PlotterList {
	if in == nil {
		return nil
	}
	out := new(PlotterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlotterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlotterSpec) DeepCopyInto(out *PlotterSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	if in.Blueprints != nil {
		in, out := &in.Blueprints, &out.Blueprints
		*out = make(map[string]BlueprintSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlotterSpec.
func (in *PlotterSpec) DeepCopy() *PlotterSpec {
	if in == nil {
		return nil
	}
	out := new(PlotterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlotterStatus) DeepCopyInto(out *PlotterStatus) {
	*out = *in
	out.ObservedState = in.ObservedState
	if in.Blueprints != nil {
		in, out := &in.Blueprints, &out.Blueprints
		*out = make(map[string]MetaBlueprint, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.ReadyTimestamp != nil {
		in, out := &in.ReadyTimestamp, &out.ReadyTimestamp
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlotterStatus.
func (in *PlotterStatus) DeepCopy() *PlotterStatus {
	if in == nil {
		return nil
	}
	out := new(PlotterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadModuleArgs) DeepCopyInto(out *ReadModuleArgs) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	if in.Transformations != nil {
		in, out := &in.Transformations, &out.Transformations
		*out = make([]protobuf.EnforcementAction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadModuleArgs.
func (in *ReadModuleArgs) DeepCopy() *ReadModuleArgs {
	if in == nil {
		return nil
	}
	out := new(ReadModuleArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceReference) DeepCopyInto(out *ResourceReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceReference.
func (in *ResourceReference) DeepCopy() *ResourceReference {
	if in == nil {
		return nil
	}
	out := new(ResourceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceStatusIndicator) DeepCopyInto(out *ResourceStatusIndicator) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceStatusIndicator.
func (in *ResourceStatusIndicator) DeepCopy() *ResourceStatusIndicator {
	if in == nil {
		return nil
	}
	out := new(ResourceStatusIndicator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WriteModuleArgs) DeepCopyInto(out *WriteModuleArgs) {
	*out = *in
	in.Destination.DeepCopyInto(&out.Destination)
	if in.Transformations != nil {
		in, out := &in.Transformations, &out.Transformations
		*out = make([]protobuf.EnforcementAction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WriteModuleArgs.
func (in *WriteModuleArgs) DeepCopy() *WriteModuleArgs {
	if in == nil {
		return nil
	}
	out := new(WriteModuleArgs)
	in.DeepCopyInto(out)
	return out
}
