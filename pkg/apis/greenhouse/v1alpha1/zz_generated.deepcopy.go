//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Greenhouse contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authentication) DeepCopyInto(out *Authentication) {
	*out = *in
	if in.OIDCConfig != nil {
		in, out := &in.OIDCConfig, &out.OIDCConfig
		*out = new(OIDCConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authentication.
func (in *Authentication) DeepCopy() *Authentication {
	if in == nil {
		return nil
	}
	out := new(Authentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	in.BearerTokenExpirationTimestamp.DeepCopyInto(&out.BearerTokenExpirationTimestamp)
	if in.HeadScaleStatus != nil {
		in, out := &in.HeadScaleStatus, &out.HeadScaleStatus
		*out = new(HeadScaleMachineStatus)
		(*in).DeepCopyInto(*out)
	}
	in.StatusConditions.DeepCopyInto(&out.StatusConditions)
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make(map[string]NodeStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
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
func (in *HeadScaleMachineStatus) DeepCopyInto(out *HeadScaleMachineStatus) {
	*out = *in
	if in.IPAddresses != nil {
		in, out := &in.IPAddresses, &out.IPAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Expiry.DeepCopyInto(&out.Expiry)
	in.CreatedAt.DeepCopyInto(&out.CreatedAt)
	if in.ForcedTags != nil {
		in, out := &in.ForcedTags, &out.ForcedTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PreAuthKey != nil {
		in, out := &in.PreAuthKey, &out.PreAuthKey
		*out = new(PreAuthKey)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeadScaleMachineStatus.
func (in *HeadScaleMachineStatus) DeepCopy() *HeadScaleMachineStatus {
	if in == nil {
		return nil
	}
	out := new(HeadScaleMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmChartReference) DeepCopyInto(out *HelmChartReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmChartReference.
func (in *HelmChartReference) DeepCopy() *HelmChartReference {
	if in == nil {
		return nil
	}
	out := new(HelmChartReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmReleaseStatus) DeepCopyInto(out *HelmReleaseStatus) {
	*out = *in
	in.FirstDeployed.DeepCopyInto(&out.FirstDeployed)
	in.LastDeployed.DeepCopyInto(&out.LastDeployed)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmReleaseStatus.
func (in *HelmReleaseStatus) DeepCopy() *HelmReleaseStatus {
	if in == nil {
		return nil
	}
	out := new(HelmReleaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeStatus) DeepCopyInto(out *NodeStatus) {
	*out = *in
	in.StatusConditions.DeepCopyInto(&out.StatusConditions)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeStatus.
func (in *NodeStatus) DeepCopy() *NodeStatus {
	if in == nil {
		return nil
	}
	out := new(NodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OIDCConfig) DeepCopyInto(out *OIDCConfig) {
	*out = *in
	out.ClientIDReference = in.ClientIDReference
	out.ClientSecretReference = in.ClientSecretReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OIDCConfig.
func (in *OIDCConfig) DeepCopy() *OIDCConfig {
	if in == nil {
		return nil
	}
	out := new(OIDCConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Organization) DeepCopyInto(out *Organization) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Organization.
func (in *Organization) DeepCopy() *Organization {
	if in == nil {
		return nil
	}
	out := new(Organization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Organization) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationList) DeepCopyInto(out *OrganizationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Organization, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationList.
func (in *OrganizationList) DeepCopy() *OrganizationList {
	if in == nil {
		return nil
	}
	out := new(OrganizationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrganizationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationSpec) DeepCopyInto(out *OrganizationSpec) {
	*out = *in
	if in.Authentication != nil {
		in, out := &in.Authentication, &out.Authentication
		*out = new(Authentication)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationSpec.
func (in *OrganizationSpec) DeepCopy() *OrganizationSpec {
	if in == nil {
		return nil
	}
	out := new(OrganizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrganizationStatus) DeepCopyInto(out *OrganizationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrganizationStatus.
func (in *OrganizationStatus) DeepCopy() *OrganizationStatus {
	if in == nil {
		return nil
	}
	out := new(OrganizationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plugin) DeepCopyInto(out *Plugin) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plugin.
func (in *Plugin) DeepCopy() *Plugin {
	if in == nil {
		return nil
	}
	out := new(Plugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Plugin) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginConfig) DeepCopyInto(out *PluginConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginConfig.
func (in *PluginConfig) DeepCopy() *PluginConfig {
	if in == nil {
		return nil
	}
	out := new(PluginConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PluginConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginConfigList) DeepCopyInto(out *PluginConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PluginConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginConfigList.
func (in *PluginConfigList) DeepCopy() *PluginConfigList {
	if in == nil {
		return nil
	}
	out := new(PluginConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PluginConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginConfigSpec) DeepCopyInto(out *PluginConfigSpec) {
	*out = *in
	if in.OptionValues != nil {
		in, out := &in.OptionValues, &out.OptionValues
		*out = make([]PluginOptionValue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginConfigSpec.
func (in *PluginConfigSpec) DeepCopy() *PluginConfigSpec {
	if in == nil {
		return nil
	}
	out := new(PluginConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginConfigStatus) DeepCopyInto(out *PluginConfigStatus) {
	*out = *in
	if in.HelmReleaseStatus != nil {
		in, out := &in.HelmReleaseStatus, &out.HelmReleaseStatus
		*out = new(HelmReleaseStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.HelmChart != nil {
		in, out := &in.HelmChart, &out.HelmChart
		*out = new(HelmChartReference)
		**out = **in
	}
	if in.UIApplication != nil {
		in, out := &in.UIApplication, &out.UIApplication
		*out = new(UIApplicationReference)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int32)
		**out = **in
	}
	if in.ExposedServices != nil {
		in, out := &in.ExposedServices, &out.ExposedServices
		*out = make(map[string]Service, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	in.StatusConditions.DeepCopyInto(&out.StatusConditions)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginConfigStatus.
func (in *PluginConfigStatus) DeepCopy() *PluginConfigStatus {
	if in == nil {
		return nil
	}
	out := new(PluginConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginList) DeepCopyInto(out *PluginList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Plugin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginList.
func (in *PluginList) DeepCopy() *PluginList {
	if in == nil {
		return nil
	}
	out := new(PluginList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PluginList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginOption) DeepCopyInto(out *PluginOption) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(v1.JSON)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginOption.
func (in *PluginOption) DeepCopy() *PluginOption {
	if in == nil {
		return nil
	}
	out := new(PluginOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginOptionValue) DeepCopyInto(out *PluginOptionValue) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(v1.JSON)
		(*in).DeepCopyInto(*out)
	}
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(ValueFromSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginOptionValue.
func (in *PluginOptionValue) DeepCopy() *PluginOptionValue {
	if in == nil {
		return nil
	}
	out := new(PluginOptionValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginSpec) DeepCopyInto(out *PluginSpec) {
	*out = *in
	if in.HelmChart != nil {
		in, out := &in.HelmChart, &out.HelmChart
		*out = new(HelmChartReference)
		**out = **in
	}
	if in.UIApplication != nil {
		in, out := &in.UIApplication, &out.UIApplication
		*out = new(UIApplicationReference)
		**out = **in
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]PluginOption, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginSpec.
func (in *PluginSpec) DeepCopy() *PluginSpec {
	if in == nil {
		return nil
	}
	out := new(PluginSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginStatus) DeepCopyInto(out *PluginStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginStatus.
func (in *PluginStatus) DeepCopy() *PluginStatus {
	if in == nil {
		return nil
	}
	out := new(PluginStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreAuthKey) DeepCopyInto(out *PreAuthKey) {
	*out = *in
	in.CreatedAt.DeepCopyInto(&out.CreatedAt)
	in.Expiration.DeepCopyInto(&out.Expiration)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreAuthKey.
func (in *PreAuthKey) DeepCopy() *PreAuthKey {
	if in == nil {
		return nil
	}
	out := new(PreAuthKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeyReference) DeepCopyInto(out *SecretKeyReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeyReference.
func (in *SecretKeyReference) DeepCopy() *SecretKeyReference {
	if in == nil {
		return nil
	}
	out := new(SecretKeyReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusConditions) DeepCopyInto(out *StatusConditions) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusConditions.
func (in *StatusConditions) DeepCopy() *StatusConditions {
	if in == nil {
		return nil
	}
	out := new(StatusConditions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Team) DeepCopyInto(out *Team) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Team.
func (in *Team) DeepCopy() *Team {
	if in == nil {
		return nil
	}
	out := new(Team)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Team) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamList) DeepCopyInto(out *TeamList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Team, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamList.
func (in *TeamList) DeepCopy() *TeamList {
	if in == nil {
		return nil
	}
	out := new(TeamList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TeamList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamMembership) DeepCopyInto(out *TeamMembership) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamMembership.
func (in *TeamMembership) DeepCopy() *TeamMembership {
	if in == nil {
		return nil
	}
	out := new(TeamMembership)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TeamMembership) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamMembershipList) DeepCopyInto(out *TeamMembershipList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TeamMembership, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamMembershipList.
func (in *TeamMembershipList) DeepCopy() *TeamMembershipList {
	if in == nil {
		return nil
	}
	out := new(TeamMembershipList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TeamMembershipList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamMembershipSpec) DeepCopyInto(out *TeamMembershipSpec) {
	*out = *in
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]User, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamMembershipSpec.
func (in *TeamMembershipSpec) DeepCopy() *TeamMembershipSpec {
	if in == nil {
		return nil
	}
	out := new(TeamMembershipSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamMembershipStatus) DeepCopyInto(out *TeamMembershipStatus) {
	*out = *in
	if in.LastSyncedTime != nil {
		in, out := &in.LastSyncedTime, &out.LastSyncedTime
		*out = (*in).DeepCopy()
	}
	if in.LastChangedTime != nil {
		in, out := &in.LastChangedTime, &out.LastChangedTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamMembershipStatus.
func (in *TeamMembershipStatus) DeepCopy() *TeamMembershipStatus {
	if in == nil {
		return nil
	}
	out := new(TeamMembershipStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamSpec) DeepCopyInto(out *TeamSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamSpec.
func (in *TeamSpec) DeepCopy() *TeamSpec {
	if in == nil {
		return nil
	}
	out := new(TeamSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamStatus) DeepCopyInto(out *TeamStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamStatus.
func (in *TeamStatus) DeepCopy() *TeamStatus {
	if in == nil {
		return nil
	}
	out := new(TeamStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UIApplicationReference) DeepCopyInto(out *UIApplicationReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UIApplicationReference.
func (in *UIApplicationReference) DeepCopy() *UIApplicationReference {
	if in == nil {
		return nil
	}
	out := new(UIApplicationReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User) DeepCopyInto(out *User) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User.
func (in *User) DeepCopy() *User {
	if in == nil {
		return nil
	}
	out := new(User)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueFromSource) DeepCopyInto(out *ValueFromSource) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(SecretKeyReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueFromSource.
func (in *ValueFromSource) DeepCopy() *ValueFromSource {
	if in == nil {
		return nil
	}
	out := new(ValueFromSource)
	in.DeepCopyInto(out)
	return out
}
