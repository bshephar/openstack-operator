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
	memcachedv1beta1 "github.com/openstack-k8s-operators/infra-operator/apis/memcached/v1beta1"
	redisv1beta1 "github.com/openstack-k8s-operators/infra-operator/apis/redis/v1beta1"
	"github.com/openstack-k8s-operators/lib-common/modules/common/condition"
	"github.com/openstack-k8s-operators/lib-common/modules/common/route"
	"github.com/openstack-k8s-operators/lib-common/modules/storage"
	apiv1beta1 "github.com/openstack-k8s-operators/mariadb-operator/api/v1beta1"
	ovn_operatorapiv1beta1 "github.com/openstack-k8s-operators/ovn-operator/api/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CeilometerSection) DeepCopyInto(out *CeilometerSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CeilometerSection.
func (in *CeilometerSection) DeepCopy() *CeilometerSection {
	if in == nil {
		return nil
	}
	out := new(CeilometerSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CinderSection) DeepCopyInto(out *CinderSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CinderSection.
func (in *CinderSection) DeepCopy() *CinderSection {
	if in == nil {
		return nil
	}
	out := new(CinderSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSMasqSection) DeepCopyInto(out *DNSMasqSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSMasqSection.
func (in *DNSMasqSection) DeepCopy() *DNSMasqSection {
	if in == nil {
		return nil
	}
	out := new(DNSMasqSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GaleraSection) DeepCopyInto(out *GaleraSection) {
	*out = *in
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make(map[string]apiv1beta1.GaleraSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GaleraSection.
func (in *GaleraSection) DeepCopy() *GaleraSection {
	if in == nil {
		return nil
	}
	out := new(GaleraSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlanceSection) DeepCopyInto(out *GlanceSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlanceSection.
func (in *GlanceSection) DeepCopy() *GlanceSection {
	if in == nil {
		return nil
	}
	out := new(GlanceSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeatSection) DeepCopyInto(out *HeatSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeatSection.
func (in *HeatSection) DeepCopy() *HeatSection {
	if in == nil {
		return nil
	}
	out := new(HeatSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizonSection) DeepCopyInto(out *HorizonSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorizonSection.
func (in *HorizonSection) DeepCopy() *HorizonSection {
	if in == nil {
		return nil
	}
	out := new(HorizonSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IronicSection) DeepCopyInto(out *IronicSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IronicSection.
func (in *IronicSection) DeepCopy() *IronicSection {
	if in == nil {
		return nil
	}
	out := new(IronicSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeystoneSection) DeepCopyInto(out *KeystoneSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	in.APIOverride.DeepCopyInto(&out.APIOverride)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeystoneSection.
func (in *KeystoneSection) DeepCopy() *KeystoneSection {
	if in == nil {
		return nil
	}
	out := new(KeystoneSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManilaSection) DeepCopyInto(out *ManilaSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManilaSection.
func (in *ManilaSection) DeepCopy() *ManilaSection {
	if in == nil {
		return nil
	}
	out := new(ManilaSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MariadbSection) DeepCopyInto(out *MariadbSection) {
	*out = *in
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make(map[string]apiv1beta1.MariaDBSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MariadbSection.
func (in *MariadbSection) DeepCopy() *MariadbSection {
	if in == nil {
		return nil
	}
	out := new(MariadbSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemcachedSection) DeepCopyInto(out *MemcachedSection) {
	*out = *in
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make(map[string]memcachedv1beta1.MemcachedSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemcachedSection.
func (in *MemcachedSection) DeepCopy() *MemcachedSection {
	if in == nil {
		return nil
	}
	out := new(MemcachedSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetalLBConfig) DeepCopyInto(out *MetalLBConfig) {
	*out = *in
	if in.LoadBalancerIPs != nil {
		in, out := &in.LoadBalancerIPs, &out.LoadBalancerIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetalLBConfig.
func (in *MetalLBConfig) DeepCopy() *MetalLBConfig {
	if in == nil {
		return nil
	}
	out := new(MetalLBConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NeutronSection) DeepCopyInto(out *NeutronSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NeutronSection.
func (in *NeutronSection) DeepCopy() *NeutronSection {
	if in == nil {
		return nil
	}
	out := new(NeutronSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NovaSection) DeepCopyInto(out *NovaSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NovaSection.
func (in *NovaSection) DeepCopy() *NovaSection {
	if in == nil {
		return nil
	}
	out := new(NovaSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OctaviaSection) DeepCopyInto(out *OctaviaSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OctaviaSection.
func (in *OctaviaSection) DeepCopy() *OctaviaSection {
	if in == nil {
		return nil
	}
	out := new(OctaviaSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackControlPlane) DeepCopyInto(out *OpenStackControlPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackControlPlane.
func (in *OpenStackControlPlane) DeepCopy() *OpenStackControlPlane {
	if in == nil {
		return nil
	}
	out := new(OpenStackControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackControlPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackControlPlaneDefaults) DeepCopyInto(out *OpenStackControlPlaneDefaults) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackControlPlaneDefaults.
func (in *OpenStackControlPlaneDefaults) DeepCopy() *OpenStackControlPlaneDefaults {
	if in == nil {
		return nil
	}
	out := new(OpenStackControlPlaneDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackControlPlaneList) DeepCopyInto(out *OpenStackControlPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackControlPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackControlPlaneList.
func (in *OpenStackControlPlaneList) DeepCopy() *OpenStackControlPlaneList {
	if in == nil {
		return nil
	}
	out := new(OpenStackControlPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackControlPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackControlPlaneSpec) DeepCopyInto(out *OpenStackControlPlaneSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.DNS.DeepCopyInto(&out.DNS)
	in.Keystone.DeepCopyInto(&out.Keystone)
	in.Placement.DeepCopyInto(&out.Placement)
	in.Glance.DeepCopyInto(&out.Glance)
	in.Cinder.DeepCopyInto(&out.Cinder)
	in.Mariadb.DeepCopyInto(&out.Mariadb)
	in.Galera.DeepCopyInto(&out.Galera)
	in.Rabbitmq.DeepCopyInto(&out.Rabbitmq)
	in.Memcached.DeepCopyInto(&out.Memcached)
	in.Ovn.DeepCopyInto(&out.Ovn)
	in.Neutron.DeepCopyInto(&out.Neutron)
	in.Nova.DeepCopyInto(&out.Nova)
	in.Heat.DeepCopyInto(&out.Heat)
	in.Ironic.DeepCopyInto(&out.Ironic)
	in.Manila.DeepCopyInto(&out.Manila)
	in.Horizon.DeepCopyInto(&out.Horizon)
	in.Ceilometer.DeepCopyInto(&out.Ceilometer)
	in.Swift.DeepCopyInto(&out.Swift)
	in.Octavia.DeepCopyInto(&out.Octavia)
	in.Redis.DeepCopyInto(&out.Redis)
	if in.ExtraMounts != nil {
		in, out := &in.ExtraMounts, &out.ExtraMounts
		*out = make([]OpenStackExtraVolMounts, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackControlPlaneSpec.
func (in *OpenStackControlPlaneSpec) DeepCopy() *OpenStackControlPlaneSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackControlPlaneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackControlPlaneStatus) DeepCopyInto(out *OpenStackControlPlaneStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackControlPlaneStatus.
func (in *OpenStackControlPlaneStatus) DeepCopy() *OpenStackControlPlaneStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackControlPlaneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackExtraVolMounts) DeepCopyInto(out *OpenStackExtraVolMounts) {
	*out = *in
	if in.VolMounts != nil {
		in, out := &in.VolMounts, &out.VolMounts
		*out = make([]storage.VolMounts, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackExtraVolMounts.
func (in *OpenStackExtraVolMounts) DeepCopy() *OpenStackExtraVolMounts {
	if in == nil {
		return nil
	}
	out := new(OpenStackExtraVolMounts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Override) DeepCopyInto(out *Override) {
	*out = *in
	if in.Route != nil {
		in, out := &in.Route, &out.Route
		*out = new(route.OverrideSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Override.
func (in *Override) DeepCopy() *Override {
	if in == nil {
		return nil
	}
	out := new(Override)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvnResources) DeepCopyInto(out *OvnResources) {
	*out = *in
	if in.OVNDBCluster != nil {
		in, out := &in.OVNDBCluster, &out.OVNDBCluster
		*out = make(map[string]ovn_operatorapiv1beta1.OVNDBClusterSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	in.OVNNorthd.DeepCopyInto(&out.OVNNorthd)
	in.OVNController.DeepCopyInto(&out.OVNController)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvnResources.
func (in *OvnResources) DeepCopy() *OvnResources {
	if in == nil {
		return nil
	}
	out := new(OvnResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvnSection) DeepCopyInto(out *OvnSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvnSection.
func (in *OvnSection) DeepCopy() *OvnSection {
	if in == nil {
		return nil
	}
	out := new(OvnSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementSection) DeepCopyInto(out *PlacementSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementSection.
func (in *PlacementSection) DeepCopy() *PlacementSection {
	if in == nil {
		return nil
	}
	out := new(PlacementSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqSection) DeepCopyInto(out *RabbitmqSection) {
	*out = *in
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make(map[string]RabbitmqTemplate, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqSection.
func (in *RabbitmqSection) DeepCopy() *RabbitmqSection {
	if in == nil {
		return nil
	}
	out := new(RabbitmqSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqTemplate) DeepCopyInto(out *RabbitmqTemplate) {
	*out = *in
	in.RabbitmqClusterSpec.DeepCopyInto(&out.RabbitmqClusterSpec)
	if in.ExternalEndpoint != nil {
		in, out := &in.ExternalEndpoint, &out.ExternalEndpoint
		*out = new(MetalLBConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqTemplate.
func (in *RabbitmqTemplate) DeepCopy() *RabbitmqTemplate {
	if in == nil {
		return nil
	}
	out := new(RabbitmqTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisSection) DeepCopyInto(out *RedisSection) {
	*out = *in
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make(map[string]redisv1beta1.RedisSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisSection.
func (in *RedisSection) DeepCopy() *RedisSection {
	if in == nil {
		return nil
	}
	out := new(RedisSection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftSection) DeepCopyInto(out *SwiftSection) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftSection.
func (in *SwiftSection) DeepCopy() *SwiftSection {
	if in == nil {
		return nil
	}
	out := new(SwiftSection)
	in.DeepCopyInto(out)
	return out
}
