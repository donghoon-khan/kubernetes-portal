package persistentvolumeclaim

import (
	"reflect"
	"testing"

	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/donghoon-khan/kubeportal/src/app/backend/api"
)

func TestGetPersistentVolumeClaimDetail(t *testing.T) {

	cases := []struct {
		persistentVolumeClaims *v1.PersistentVolumeClaim
		expected               *PersistentVolumeClaimDetail
	}{
		{
			&v1.PersistentVolumeClaim{
				TypeMeta:   metaV1.TypeMeta{Kind: "persistentvolumeclaim"},
				ObjectMeta: metaV1.ObjectMeta{Name: "foo", Namespace: "bar"},
				Spec: v1.PersistentVolumeClaimSpec{
					AccessModes: []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce},
					Resources:   v1.ResourceRequirements{},
					VolumeName:  "volume",
				},
				Status: v1.PersistentVolumeClaimStatus{
					Phase:       v1.PersistentVolumeClaimPhase(v1.ClaimPending),
					AccessModes: []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce},
					Capacity:    nil,
				},
			},
			&PersistentVolumeClaimDetail{
				PersistentVolumeClaim: PersistentVolumeClaim{
					ObjectMeta:  api.ObjectMeta{Name: "foo", Namespace: "bar"},
					TypeMeta:    api.TypeMeta{Kind: "persistentvolumeclaim"},
					Status:      string(v1.ClaimPending),
					Volume:      "volume",
					Capacity:    nil,
					AccessModes: []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce},
				},
			},
		},
	}
	for _, c := range cases {
		actual := getPersistentVolumeClaimDetail(*c.persistentVolumeClaims)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("getPersistentVolumeClaimDetail(%#v) == \n%#v\nexpected \n%#v\n",
				c.persistentVolumeClaims, actual, c.expected)
		}
	}
}