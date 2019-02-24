package v1stable

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type Crontab struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec CrontabSpec `json:"spec"`
}

type CrontabSpec struct {
	CronSpec string `json:"cronSpec"`
}

type CrontabList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Crontab `json:"items"`
}

func (in *Crontab) DeepCopyInto(out *Crontab) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = CrontabSpec{
		CronSpec: in.Spec.CronSpec,
	}
}

func (in *Crontab) DeepCopyObject() runtime.Object {
	out := Crontab{}
	in.DeepCopyInto(&out)
	return &out
}

func (in *CrontabList) DeepCopyObject() runtime.Object {
	out := CrontabList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]Crontab, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}
	return &out
}
