package resources

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

const (
	deploymentName = "people-backend"
)

var (
	deploymentSelector = v1.SetAsLabelSelector(labels.Set{
		"app": deploymentName,
	})
)

func NewDeployment(namespace, image string, replicas int32) *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: v1.ObjectMeta{
			Namespace:   namespace,
			Name:        deploymentName,
			Annotations: map[string]string{},
			Labels:      map[string]string{},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: deploymentSelector,
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  deploymentName,
							Image: image,
						},
					},
				},
			},
		},
	}
}
