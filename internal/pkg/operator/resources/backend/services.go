package backend

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	serviceName = "people-backend-service"
)

var (
	serviceSelectorLabels = deploymentSelectorLabels
)

func NewService(namespace string, httpPort int32) *v1.Service {
	return &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:   namespace,
			Name:        serviceName,
			Annotations: map[string]string{},
			Labels:      map[string]string{},
		},
		Spec: v1.ServiceSpec{
			Ports: []v1.ServicePort{{
				Name:       "http",
				TargetPort: intstr.FromInt(int(httpPort)),
				Port:       httpPort,
			}},
			Selector: serviceSelectorLabels,
			Type:     v1.ServiceTypeNodePort,
		},
	}
}
