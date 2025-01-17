// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppProtect) DeepCopyInto(out *AppProtect) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppProtect.
func (in *AppProtect) DeepCopy() *AppProtect {
	if in == nil {
		return nil
	}
	out := new(AppProtect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthStatus) DeepCopyInto(out *HealthStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthStatus.
func (in *HealthStatus) DeepCopy() *HealthStatus {
	if in == nil {
		return nil
	}
	out := new(HealthStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NginxIngressController) DeepCopyInto(out *NginxIngressController) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NginxIngressController.
func (in *NginxIngressController) DeepCopy() *NginxIngressController {
	if in == nil {
		return nil
	}
	out := new(NginxIngressController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NginxIngressController) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NginxIngressControllerList) DeepCopyInto(out *NginxIngressControllerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NginxIngressController, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NginxIngressControllerList.
func (in *NginxIngressControllerList) DeepCopy() *NginxIngressControllerList {
	if in == nil {
		return nil
	}
	out := new(NginxIngressControllerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NginxIngressControllerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NginxIngressControllerSpec) DeepCopyInto(out *NginxIngressControllerSpec) {
	*out = *in
	out.Image = in.Image
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.EnableCRDs != nil {
		in, out := &in.EnableCRDs, &out.EnableCRDs
		*out = new(bool)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(Service)
		(*in).DeepCopyInto(*out)
	}
	if in.Workload != nil {
		in, out := &in.Workload, &out.Workload
		*out = new(Workload)
		(*in).DeepCopyInto(*out)
	}
	if in.HealthStatus != nil {
		in, out := &in.HealthStatus, &out.HealthStatus
		*out = new(HealthStatus)
		**out = **in
	}
	if in.NginxStatus != nil {
		in, out := &in.NginxStatus, &out.NginxStatus
		*out = new(NginxStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.ReportIngressStatus != nil {
		in, out := &in.ReportIngressStatus, &out.ReportIngressStatus
		*out = new(ReportIngressStatus)
		**out = **in
	}
	if in.EnableLeaderElection != nil {
		in, out := &in.EnableLeaderElection, &out.EnableLeaderElection
		*out = new(bool)
		**out = **in
	}
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(Prometheus)
		(*in).DeepCopyInto(*out)
	}
	if in.ConfigMapData != nil {
		in, out := &in.ConfigMapData, &out.ConfigMapData
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AppProtect != nil {
		in, out := &in.AppProtect, &out.AppProtect
		*out = new(AppProtect)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NginxIngressControllerSpec.
func (in *NginxIngressControllerSpec) DeepCopy() *NginxIngressControllerSpec {
	if in == nil {
		return nil
	}
	out := new(NginxIngressControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NginxIngressControllerStatus) DeepCopyInto(out *NginxIngressControllerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NginxIngressControllerStatus.
func (in *NginxIngressControllerStatus) DeepCopy() *NginxIngressControllerStatus {
	if in == nil {
		return nil
	}
	out := new(NginxIngressControllerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NginxStatus) DeepCopyInto(out *NginxStatus) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(uint16)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NginxStatus.
func (in *NginxStatus) DeepCopy() *NginxStatus {
	if in == nil {
		return nil
	}
	out := new(NginxStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Prometheus) DeepCopyInto(out *Prometheus) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(uint16)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Prometheus.
func (in *Prometheus) DeepCopy() *Prometheus {
	if in == nil {
		return nil
	}
	out := new(Prometheus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportIngressStatus) DeepCopyInto(out *ReportIngressStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportIngressStatus.
func (in *ReportIngressStatus) DeepCopy() *ReportIngressStatus {
	if in == nil {
		return nil
	}
	out := new(ReportIngressStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	if in.ExtraLabels != nil {
		in, out := &in.ExtraLabels, &out.ExtraLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
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
func (in *Workload) DeepCopyInto(out *Workload) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	if in.ExtraLabels != nil {
		in, out := &in.ExtraLabels, &out.ExtraLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workload.
func (in *Workload) DeepCopy() *Workload {
	if in == nil {
		return nil
	}
	out := new(Workload)
	in.DeepCopyInto(out)
	return out
}
