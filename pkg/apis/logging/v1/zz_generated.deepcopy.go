// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1

import (
	outputs "github.com/openshift/cluster-logging-operator/pkg/apis/logging/v1/outputs"
	loggingv1 "github.com/openshift/elasticsearch-operator/pkg/apis/logging/v1"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCondition) DeepCopyInto(out *ClusterCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCondition.
func (in *ClusterCondition) DeepCopy() *ClusterCondition {
	if in == nil {
		return nil
	}
	out := new(ClusterCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ClusterConditions) DeepCopyInto(out *ClusterConditions) {
	{
		in := &in
		*out = make(ClusterConditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConditions.
func (in ClusterConditions) DeepCopy() ClusterConditions {
	if in == nil {
		return nil
	}
	out := new(ClusterConditions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterLogForwarder) DeepCopyInto(out *ClusterLogForwarder) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterLogForwarder.
func (in *ClusterLogForwarder) DeepCopy() *ClusterLogForwarder {
	if in == nil {
		return nil
	}
	out := new(ClusterLogForwarder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterLogForwarder) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterLogForwarderList) DeepCopyInto(out *ClusterLogForwarderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterLogForwarder, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterLogForwarderList.
func (in *ClusterLogForwarderList) DeepCopy() *ClusterLogForwarderList {
	if in == nil {
		return nil
	}
	out := new(ClusterLogForwarderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterLogForwarderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterLogForwarderSpec) DeepCopyInto(out *ClusterLogForwarderSpec) {
	*out = *in
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = make([]InputSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = make([]OutputSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Pipelines != nil {
		in, out := &in.Pipelines, &out.Pipelines
		*out = make([]PipelineSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterLogForwarderSpec.
func (in *ClusterLogForwarderSpec) DeepCopy() *ClusterLogForwarderSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterLogForwarderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterLogForwarderStatus) DeepCopyInto(out *ClusterLogForwarderStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(Conditions, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Inputs != nil {
		in, out := &in.Inputs, &out.Inputs
		*out = make(NamedConditions, len(*in))
		for key, val := range *in {
			var outVal map[ConditionType]Condition
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(Conditions, len(*in))
				for key, val := range *in {
					(*out)[key] = *val.DeepCopy()
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = make(NamedConditions, len(*in))
		for key, val := range *in {
			var outVal map[ConditionType]Condition
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(Conditions, len(*in))
				for key, val := range *in {
					(*out)[key] = *val.DeepCopy()
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Pipelines != nil {
		in, out := &in.Pipelines, &out.Pipelines
		*out = make(NamedConditions, len(*in))
		for key, val := range *in {
			var outVal map[ConditionType]Condition
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(Conditions, len(*in))
				for key, val := range *in {
					(*out)[key] = *val.DeepCopy()
				}
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterLogForwarderStatus.
func (in *ClusterLogForwarderStatus) DeepCopy() *ClusterLogForwarderStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterLogForwarderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterLogging) DeepCopyInto(out *ClusterLogging) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterLogging.
func (in *ClusterLogging) DeepCopy() *ClusterLogging {
	if in == nil {
		return nil
	}
	out := new(ClusterLogging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterLogging) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterLoggingList) DeepCopyInto(out *ClusterLoggingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterLogging, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterLoggingList.
func (in *ClusterLoggingList) DeepCopy() *ClusterLoggingList {
	if in == nil {
		return nil
	}
	out := new(ClusterLoggingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterLoggingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterLoggingSpec) DeepCopyInto(out *ClusterLoggingSpec) {
	*out = *in
	if in.Visualization != nil {
		in, out := &in.Visualization, &out.Visualization
		*out = new(VisualizationSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.LogStore != nil {
		in, out := &in.LogStore, &out.LogStore
		*out = new(LogStoreSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Collection != nil {
		in, out := &in.Collection, &out.Collection
		*out = new(CollectionSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Curation != nil {
		in, out := &in.Curation, &out.Curation
		*out = new(CurationSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterLoggingSpec.
func (in *ClusterLoggingSpec) DeepCopy() *ClusterLoggingSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterLoggingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterLoggingStatus) DeepCopyInto(out *ClusterLoggingStatus) {
	*out = *in
	in.Visualization.DeepCopyInto(&out.Visualization)
	in.LogStore.DeepCopyInto(&out.LogStore)
	in.Collection.DeepCopyInto(&out.Collection)
	in.Curation.DeepCopyInto(&out.Curation)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ClusterCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterLoggingStatus.
func (in *ClusterLoggingStatus) DeepCopy() *ClusterLoggingStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterLoggingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectionSpec) DeepCopyInto(out *CollectionSpec) {
	*out = *in
	in.Logs.DeepCopyInto(&out.Logs)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectionSpec.
func (in *CollectionSpec) DeepCopy() *CollectionSpec {
	if in == nil {
		return nil
	}
	out := new(CollectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectionStatus) DeepCopyInto(out *CollectionStatus) {
	*out = *in
	in.Logs.DeepCopyInto(&out.Logs)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectionStatus.
func (in *CollectionStatus) DeepCopy() *CollectionStatus {
	if in == nil {
		return nil
	}
	out := new(CollectionStatus)
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
func (in Conditions) DeepCopyInto(out *Conditions) {
	{
		in := &in
		*out = make(Conditions, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Conditions.
func (in Conditions) DeepCopy() Conditions {
	if in == nil {
		return nil
	}
	out := new(Conditions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CurationSpec) DeepCopyInto(out *CurationSpec) {
	*out = *in
	in.CuratorSpec.DeepCopyInto(&out.CuratorSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CurationSpec.
func (in *CurationSpec) DeepCopy() *CurationSpec {
	if in == nil {
		return nil
	}
	out := new(CurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CurationStatus) DeepCopyInto(out *CurationStatus) {
	*out = *in
	if in.CuratorStatus != nil {
		in, out := &in.CuratorStatus, &out.CuratorStatus
		*out = make([]CuratorStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CurationStatus.
func (in *CurationStatus) DeepCopy() *CurationStatus {
	if in == nil {
		return nil
	}
	out := new(CurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CuratorSpec) DeepCopyInto(out *CuratorSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CuratorSpec.
func (in *CuratorSpec) DeepCopy() *CuratorSpec {
	if in == nil {
		return nil
	}
	out := new(CuratorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CuratorStatus) DeepCopyInto(out *CuratorStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(map[string]ClusterConditions, len(*in))
		for key, val := range *in {
			var outVal []ClusterCondition
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(ClusterConditions, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CuratorStatus.
func (in *CuratorStatus) DeepCopy() *CuratorStatus {
	if in == nil {
		return nil
	}
	out := new(CuratorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ElasticsearchClusterConditions) DeepCopyInto(out *ElasticsearchClusterConditions) {
	{
		in := &in
		*out = make(ElasticsearchClusterConditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchClusterConditions.
func (in ElasticsearchClusterConditions) DeepCopy() ElasticsearchClusterConditions {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchClusterConditions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchSpec) DeepCopyInto(out *ElasticsearchSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Storage.DeepCopyInto(&out.Storage)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchSpec.
func (in *ElasticsearchSpec) DeepCopy() *ElasticsearchSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticsearchStatus) DeepCopyInto(out *ElasticsearchStatus) {
	*out = *in
	if in.ReplicaSets != nil {
		in, out := &in.ReplicaSets, &out.ReplicaSets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Deployments != nil {
		in, out := &in.Deployments, &out.Deployments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.StatefulSets != nil {
		in, out := &in.StatefulSets, &out.StatefulSets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Cluster = in.Cluster
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make(map[ElasticsearchRoleType]PodStateMap, len(*in))
		for key, val := range *in {
			var outVal map[PodStateType][]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(PodStateMap, len(*in))
				for key, val := range *in {
					var outVal []string
					if val == nil {
						(*out)[key] = nil
					} else {
						in, out := &val, &outVal
						*out = make([]string, len(*in))
						copy(*out, *in)
					}
					(*out)[key] = outVal
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.ClusterConditions != nil {
		in, out := &in.ClusterConditions, &out.ClusterConditions
		*out = make(ElasticsearchClusterConditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeConditions != nil {
		in, out := &in.NodeConditions, &out.NodeConditions
		*out = make(map[string]ElasticsearchClusterConditions, len(*in))
		for key, val := range *in {
			var outVal []loggingv1.ClusterCondition
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(ElasticsearchClusterConditions, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticsearchStatus.
func (in *ElasticsearchStatus) DeepCopy() *ElasticsearchStatus {
	if in == nil {
		return nil
	}
	out := new(ElasticsearchStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventCollectionSpec) DeepCopyInto(out *EventCollectionSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventCollectionSpec.
func (in *EventCollectionSpec) DeepCopy() *EventCollectionSpec {
	if in == nil {
		return nil
	}
	out := new(EventCollectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventCollectionStatus) DeepCopyInto(out *EventCollectionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventCollectionStatus.
func (in *EventCollectionStatus) DeepCopy() *EventCollectionStatus {
	if in == nil {
		return nil
	}
	out := new(EventCollectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FluentdCollectorStatus) DeepCopyInto(out *FluentdCollectorStatus) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make(PodStateMap, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(map[string]ClusterConditions, len(*in))
		for key, val := range *in {
			var outVal []ClusterCondition
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(ClusterConditions, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FluentdCollectorStatus.
func (in *FluentdCollectorStatus) DeepCopy() *FluentdCollectorStatus {
	if in == nil {
		return nil
	}
	out := new(FluentdCollectorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FluentdNormalizerStatus) DeepCopyInto(out *FluentdNormalizerStatus) {
	*out = *in
	if in.ReplicaSets != nil {
		in, out := &in.ReplicaSets, &out.ReplicaSets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make(PodStateMap, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(map[string]ClusterConditions, len(*in))
		for key, val := range *in {
			var outVal []ClusterCondition
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(ClusterConditions, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FluentdNormalizerStatus.
func (in *FluentdNormalizerStatus) DeepCopy() *FluentdNormalizerStatus {
	if in == nil {
		return nil
	}
	out := new(FluentdNormalizerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FluentdSpec) DeepCopyInto(out *FluentdSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FluentdSpec.
func (in *FluentdSpec) DeepCopy() *FluentdSpec {
	if in == nil {
		return nil
	}
	out := new(FluentdSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InputSpec) DeepCopyInto(out *InputSpec) {
	*out = *in
	in.InputTypeSpec.DeepCopyInto(&out.InputTypeSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InputSpec.
func (in *InputSpec) DeepCopy() *InputSpec {
	if in == nil {
		return nil
	}
	out := new(InputSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InputTypeSpec) DeepCopyInto(out *InputTypeSpec) {
	*out = *in
	if in.Application != nil {
		in, out := &in.Application, &out.Application
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InputTypeSpec.
func (in *InputTypeSpec) DeepCopy() *InputTypeSpec {
	if in == nil {
		return nil
	}
	out := new(InputTypeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KibanaSpec) DeepCopyInto(out *KibanaSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ProxySpec.DeepCopyInto(&out.ProxySpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KibanaSpec.
func (in *KibanaSpec) DeepCopy() *KibanaSpec {
	if in == nil {
		return nil
	}
	out := new(KibanaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KibanaStatus) DeepCopyInto(out *KibanaStatus) {
	*out = *in
	if in.ReplicaSets != nil {
		in, out := &in.ReplicaSets, &out.ReplicaSets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make(PodStateMap, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(map[string]ClusterConditions, len(*in))
		for key, val := range *in {
			var outVal []ClusterCondition
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(ClusterConditions, len(*in))
				for i := range *in {
					(*in)[i].DeepCopyInto(&(*out)[i])
				}
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KibanaStatus.
func (in *KibanaStatus) DeepCopy() *KibanaStatus {
	if in == nil {
		return nil
	}
	out := new(KibanaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogCollectionSpec) DeepCopyInto(out *LogCollectionSpec) {
	*out = *in
	in.FluentdSpec.DeepCopyInto(&out.FluentdSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogCollectionSpec.
func (in *LogCollectionSpec) DeepCopy() *LogCollectionSpec {
	if in == nil {
		return nil
	}
	out := new(LogCollectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogCollectionStatus) DeepCopyInto(out *LogCollectionStatus) {
	*out = *in
	in.FluentdStatus.DeepCopyInto(&out.FluentdStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogCollectionStatus.
func (in *LogCollectionStatus) DeepCopy() *LogCollectionStatus {
	if in == nil {
		return nil
	}
	out := new(LogCollectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogStoreSpec) DeepCopyInto(out *LogStoreSpec) {
	*out = *in
	in.ElasticsearchSpec.DeepCopyInto(&out.ElasticsearchSpec)
	if in.RetentionPolicy != nil {
		in, out := &in.RetentionPolicy, &out.RetentionPolicy
		*out = new(RetentionPoliciesSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogStoreSpec.
func (in *LogStoreSpec) DeepCopy() *LogStoreSpec {
	if in == nil {
		return nil
	}
	out := new(LogStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogStoreStatus) DeepCopyInto(out *LogStoreStatus) {
	*out = *in
	if in.ElasticsearchStatus != nil {
		in, out := &in.ElasticsearchStatus, &out.ElasticsearchStatus
		*out = make([]ElasticsearchStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogStoreStatus.
func (in *LogStoreStatus) DeepCopy() *LogStoreStatus {
	if in == nil {
		return nil
	}
	out := new(LogStoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in NamedConditions) DeepCopyInto(out *NamedConditions) {
	{
		in := &in
		*out = make(NamedConditions, len(*in))
		for key, val := range *in {
			var outVal map[ConditionType]Condition
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(Conditions, len(*in))
				for key, val := range *in {
					(*out)[key] = *val.DeepCopy()
				}
			}
			(*out)[key] = outVal
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedConditions.
func (in NamedConditions) DeepCopy() NamedConditions {
	if in == nil {
		return nil
	}
	out := new(NamedConditions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NormalizerStatus) DeepCopyInto(out *NormalizerStatus) {
	*out = *in
	if in.FluentdStatus != nil {
		in, out := &in.FluentdStatus, &out.FluentdStatus
		*out = make([]FluentdNormalizerStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NormalizerStatus.
func (in *NormalizerStatus) DeepCopy() *NormalizerStatus {
	if in == nil {
		return nil
	}
	out := new(NormalizerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutputSecretSpec) DeepCopyInto(out *OutputSecretSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutputSecretSpec.
func (in *OutputSecretSpec) DeepCopy() *OutputSecretSpec {
	if in == nil {
		return nil
	}
	out := new(OutputSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutputSpec) DeepCopyInto(out *OutputSpec) {
	*out = *in
	in.OutputTypeSpec.DeepCopyInto(&out.OutputTypeSpec)
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(OutputSecretSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutputSpec.
func (in *OutputSpec) DeepCopy() *OutputSpec {
	if in == nil {
		return nil
	}
	out := new(OutputSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutputTypeSpec) DeepCopyInto(out *OutputTypeSpec) {
	*out = *in
	if in.Syslog != nil {
		in, out := &in.Syslog, &out.Syslog
		*out = new(outputs.Syslog)
		**out = **in
	}
	if in.FluentForward != nil {
		in, out := &in.FluentForward, &out.FluentForward
		*out = new(outputs.FluentForward)
		**out = **in
	}
	if in.ElasticSearch != nil {
		in, out := &in.ElasticSearch, &out.ElasticSearch
		*out = new(outputs.ElasticSearch)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutputTypeSpec.
func (in *OutputTypeSpec) DeepCopy() *OutputTypeSpec {
	if in == nil {
		return nil
	}
	out := new(OutputTypeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpec) DeepCopyInto(out *PipelineSpec) {
	*out = *in
	if in.OutputRefs != nil {
		in, out := &in.OutputRefs, &out.OutputRefs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.InputRefs != nil {
		in, out := &in.InputRefs, &out.InputRefs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpec.
func (in *PipelineSpec) DeepCopy() *PipelineSpec {
	if in == nil {
		return nil
	}
	out := new(PipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PodStateMap) DeepCopyInto(out *PodStateMap) {
	{
		in := &in
		*out = make(PodStateMap, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodStateMap.
func (in PodStateMap) DeepCopy() PodStateMap {
	if in == nil {
		return nil
	}
	out := new(PodStateMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxySpec) DeepCopyInto(out *ProxySpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxySpec.
func (in *ProxySpec) DeepCopy() *ProxySpec {
	if in == nil {
		return nil
	}
	out := new(ProxySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Reconnect) DeepCopyInto(out *Reconnect) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Reconnect.
func (in *Reconnect) DeepCopy() *Reconnect {
	if in == nil {
		return nil
	}
	out := new(Reconnect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionPoliciesSpec) DeepCopyInto(out *RetentionPoliciesSpec) {
	*out = *in
	if in.App != nil {
		in, out := &in.App, &out.App
		*out = new(RetentionPolicySpec)
		**out = **in
	}
	if in.Infra != nil {
		in, out := &in.Infra, &out.Infra
		*out = new(RetentionPolicySpec)
		**out = **in
	}
	if in.Audit != nil {
		in, out := &in.Audit, &out.Audit
		*out = new(RetentionPolicySpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionPoliciesSpec.
func (in *RetentionPoliciesSpec) DeepCopy() *RetentionPoliciesSpec {
	if in == nil {
		return nil
	}
	out := new(RetentionPoliciesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionPolicySpec) DeepCopyInto(out *RetentionPolicySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionPolicySpec.
func (in *RetentionPolicySpec) DeepCopy() *RetentionPolicySpec {
	if in == nil {
		return nil
	}
	out := new(RetentionPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in RouteMap) DeepCopyInto(out *RouteMap) {
	{
		in := &in
		*out = make(RouteMap, len(*in))
		for key, val := range *in {
			var outVal map[string]sets.Empty
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(sets.String, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteMap.
func (in RouteMap) DeepCopy() RouteMap {
	if in == nil {
		return nil
	}
	out := new(RouteMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Routes) DeepCopyInto(out *Routes) {
	*out = *in
	if in.ByInput != nil {
		in, out := &in.ByInput, &out.ByInput
		*out = make(RouteMap, len(*in))
		for key, val := range *in {
			var outVal map[string]sets.Empty
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(sets.String, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.ByOutput != nil {
		in, out := &in.ByOutput, &out.ByOutput
		*out = make(RouteMap, len(*in))
		for key, val := range *in {
			var outVal map[string]sets.Empty
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(sets.String, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Routes.
func (in *Routes) DeepCopy() *Routes {
	if in == nil {
		return nil
	}
	out := new(Routes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VisualizationSpec) DeepCopyInto(out *VisualizationSpec) {
	*out = *in
	in.KibanaSpec.DeepCopyInto(&out.KibanaSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VisualizationSpec.
func (in *VisualizationSpec) DeepCopy() *VisualizationSpec {
	if in == nil {
		return nil
	}
	out := new(VisualizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VisualizationStatus) DeepCopyInto(out *VisualizationStatus) {
	*out = *in
	if in.KibanaStatus != nil {
		in, out := &in.KibanaStatus, &out.KibanaStatus
		*out = make([]loggingv1.KibanaStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VisualizationStatus.
func (in *VisualizationStatus) DeepCopy() *VisualizationStatus {
	if in == nil {
		return nil
	}
	out := new(VisualizationStatus)
	in.DeepCopyInto(out)
	return out
}
