// +build !ignore_autogenerated

/*
Copyright 2019 Red Hat Inc.

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
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStorageConfig) DeepCopyInto(out *BackupStorageConfig) {
	*out = *in
	if in.CredsSecretRef != nil {
		in, out := &in.CredsSecretRef, &out.CredsSecretRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStorageConfig.
func (in *BackupStorageConfig) DeepCopy() *BackupStorageConfig {
	if in == nil {
		return nil
	}
	out := new(BackupStorageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
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
func (in *Conditions) DeepCopyInto(out *Conditions) {
	*out = *in
	if in.List != nil {
		in, out := &in.List, &out.List
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Conditions.
func (in *Conditions) DeepCopy() *Conditions {
	if in == nil {
		return nil
	}
	out := new(Conditions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigCluster) DeepCopyInto(out *MigCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigCluster.
func (in *MigCluster) DeepCopy() *MigCluster {
	if in == nil {
		return nil
	}
	out := new(MigCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigClusterList) DeepCopyInto(out *MigClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MigCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigClusterList.
func (in *MigClusterList) DeepCopy() *MigClusterList {
	if in == nil {
		return nil
	}
	out := new(MigClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigClusterSpec) DeepCopyInto(out *MigClusterSpec) {
	*out = *in
	if in.ClusterRef != nil {
		in, out := &in.ClusterRef, &out.ClusterRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.ServiceAccountSecretRef != nil {
		in, out := &in.ServiceAccountSecretRef, &out.ServiceAccountSecretRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigClusterSpec.
func (in *MigClusterSpec) DeepCopy() *MigClusterSpec {
	if in == nil {
		return nil
	}
	out := new(MigClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigClusterStatus) DeepCopyInto(out *MigClusterStatus) {
	*out = *in
	in.Conditions.DeepCopyInto(&out.Conditions)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigClusterStatus.
func (in *MigClusterStatus) DeepCopy() *MigClusterStatus {
	if in == nil {
		return nil
	}
	out := new(MigClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigMigration) DeepCopyInto(out *MigMigration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigMigration.
func (in *MigMigration) DeepCopy() *MigMigration {
	if in == nil {
		return nil
	}
	out := new(MigMigration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigMigration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigMigrationList) DeepCopyInto(out *MigMigrationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MigMigration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigMigrationList.
func (in *MigMigrationList) DeepCopy() *MigMigrationList {
	if in == nil {
		return nil
	}
	out := new(MigMigrationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigMigrationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigMigrationSpec) DeepCopyInto(out *MigMigrationSpec) {
	*out = *in
	if in.MigPlanRef != nil {
		in, out := &in.MigPlanRef, &out.MigPlanRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigMigrationSpec.
func (in *MigMigrationSpec) DeepCopy() *MigMigrationSpec {
	if in == nil {
		return nil
	}
	out := new(MigMigrationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigMigrationStatus) DeepCopyInto(out *MigMigrationStatus) {
	*out = *in
	in.Conditions.DeepCopyInto(&out.Conditions)
	if in.StartTimestamp != nil {
		in, out := &in.StartTimestamp, &out.StartTimestamp
		*out = (*in).DeepCopy()
	}
	if in.CompletionTimestamp != nil {
		in, out := &in.CompletionTimestamp, &out.CompletionTimestamp
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigMigrationStatus.
func (in *MigMigrationStatus) DeepCopy() *MigMigrationStatus {
	if in == nil {
		return nil
	}
	out := new(MigMigrationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigPlan) DeepCopyInto(out *MigPlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigPlan.
func (in *MigPlan) DeepCopy() *MigPlan {
	if in == nil {
		return nil
	}
	out := new(MigPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigPlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigPlanList) DeepCopyInto(out *MigPlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MigPlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigPlanList.
func (in *MigPlanList) DeepCopy() *MigPlanList {
	if in == nil {
		return nil
	}
	out := new(MigPlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigPlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigPlanSpec) DeepCopyInto(out *MigPlanSpec) {
	*out = *in
	in.PersistentVolumes.DeepCopyInto(&out.PersistentVolumes)
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SrcMigClusterRef != nil {
		in, out := &in.SrcMigClusterRef, &out.SrcMigClusterRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.DestMigClusterRef != nil {
		in, out := &in.DestMigClusterRef, &out.DestMigClusterRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.MigStorageRef != nil {
		in, out := &in.MigStorageRef, &out.MigStorageRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigPlanSpec.
func (in *MigPlanSpec) DeepCopy() *MigPlanSpec {
	if in == nil {
		return nil
	}
	out := new(MigPlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigPlanStatus) DeepCopyInto(out *MigPlanStatus) {
	*out = *in
	in.Conditions.DeepCopyInto(&out.Conditions)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigPlanStatus.
func (in *MigPlanStatus) DeepCopy() *MigPlanStatus {
	if in == nil {
		return nil
	}
	out := new(MigPlanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigStorage) DeepCopyInto(out *MigStorage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigStorage.
func (in *MigStorage) DeepCopy() *MigStorage {
	if in == nil {
		return nil
	}
	out := new(MigStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigStorage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigStorageList) DeepCopyInto(out *MigStorageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MigStorage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigStorageList.
func (in *MigStorageList) DeepCopy() *MigStorageList {
	if in == nil {
		return nil
	}
	out := new(MigStorageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigStorageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigStorageSpec) DeepCopyInto(out *MigStorageSpec) {
	*out = *in
	in.BackupStorageConfig.DeepCopyInto(&out.BackupStorageConfig)
	in.VolumeSnapshotConfig.DeepCopyInto(&out.VolumeSnapshotConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigStorageSpec.
func (in *MigStorageSpec) DeepCopy() *MigStorageSpec {
	if in == nil {
		return nil
	}
	out := new(MigStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigStorageStatus) DeepCopyInto(out *MigStorageStatus) {
	*out = *in
	in.Conditions.DeepCopyInto(&out.Conditions)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigStorageStatus.
func (in *MigStorageStatus) DeepCopy() *MigStorageStatus {
	if in == nil {
		return nil
	}
	out := new(MigStorageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolume) DeepCopyInto(out *PersistentVolume) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolume.
func (in *PersistentVolume) DeepCopy() *PersistentVolume {
	if in == nil {
		return nil
	}
	out := new(PersistentVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumes) DeepCopyInto(out *PersistentVolumes) {
	*out = *in
	if in.List != nil {
		in, out := &in.List, &out.List
		*out = make([]PersistentVolume, len(*in))
		copy(*out, *in)
	}
	if in.index != nil {
		in, out := &in.index, &out.index
		*out = make(map[string]int, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumes.
func (in *PersistentVolumes) DeepCopy() *PersistentVolumes {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanResources) DeepCopyInto(out *PlanResources) {
	*out = *in
	if in.MigPlan != nil {
		in, out := &in.MigPlan, &out.MigPlan
		*out = new(MigPlan)
		(*in).DeepCopyInto(*out)
	}
	if in.MigStorage != nil {
		in, out := &in.MigStorage, &out.MigStorage
		*out = new(MigStorage)
		(*in).DeepCopyInto(*out)
	}
	if in.SrcMigCluster != nil {
		in, out := &in.SrcMigCluster, &out.SrcMigCluster
		*out = new(MigCluster)
		(*in).DeepCopyInto(*out)
	}
	if in.DestMigCluster != nil {
		in, out := &in.DestMigCluster, &out.DestMigCluster
		*out = new(MigCluster)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanResources.
func (in *PlanResources) DeepCopy() *PlanResources {
	if in == nil {
		return nil
	}
	out := new(PlanResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSnapshotConfig) DeepCopyInto(out *VolumeSnapshotConfig) {
	*out = *in
	if in.CredsSecretRef != nil {
		in, out := &in.CredsSecretRef, &out.CredsSecretRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSnapshotConfig.
func (in *VolumeSnapshotConfig) DeepCopy() *VolumeSnapshotConfig {
	if in == nil {
		return nil
	}
	out := new(VolumeSnapshotConfig)
	in.DeepCopyInto(out)
	return out
}
