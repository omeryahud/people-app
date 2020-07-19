package database

import (
	"strconv"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

const (
	// Deployment constants
	deploymentName = "people-database"

	// Environment constants
	httpPortKey = "HTTP_PORT"
)

var (
	deploymentSelectorLabels = labels.Set{
		"app": deploymentName,
	}

	deploymentSelector = v1.SetAsLabelSelector(deploymentSelectorLabels)
)

func NewDeployment(namespace, image string, httpPort int32, replicas int32) *appsv1.Deployment {
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
				ObjectMeta: v1.ObjectMeta{
					Labels: deploymentSelectorLabels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            deploymentName,
							Image:           image,
							ImagePullPolicy: corev1.PullAlways,
							Ports: []corev1.ContainerPort{{
								Name:          "http",
								ContainerPort: httpPort,
							}},
							Env: []corev1.EnvVar{{
								Name:  httpPortKey,
								Value: strconv.Itoa(int(httpPort)),
							}},
						},
					},
				},
			},
		},
	}
}
