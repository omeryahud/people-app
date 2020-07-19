package reconcile

import (
	"context"
	"encoding/json"

	peoplev1alpha1 "github.com/omeryahud/people-app/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	clnt "sigs.k8s.io/controller-runtime/pkg/client"
)

func createOrPatchDeployment(client client.Client,
	ctx context.Context,
	deployment *appsv1.Deployment) error {
	err := client.Create(ctx, deployment)

	if err != nil {
		if errors.IsAlreadyExists(err) {
			patchData, err := json.Marshal(deployment)
			if err != nil {
				return err
			}

			err = client.Patch(ctx,
				deployment,
				clnt.ConstantPatch(types.MergePatchType, patchData))
			if err != nil {
				return err
			}

			return nil
		} else {
			return err
		}
	}

	return nil
}

func createOrPatchService(client client.Client,
	ctx context.Context,
	service *v1.Service) error {
	err := client.Create(ctx, service)

	if err != nil {
		if errors.IsAlreadyExists(err) {
			patchData, err := json.Marshal(service)
			if err != nil {
				return nil
			}

			err = client.Patch(ctx,
				service,
				clnt.ConstantPatch(types.StrategicMergePatchType, patchData))
			if err != nil {
				return err
			}

			return nil
		} else {
			return err
		}
	}

	return nil
}

func injectOwnerReference(owner *peoplev1alpha1.PeopleApp, object *metav1.ObjectMeta) {
	object.SetOwnerReferences([]metav1.OwnerReference{{
		APIVersion: owner.APIVersion,
		Controller: func() *bool {
			b := true
			return &b
		}(),
		Kind: owner.Kind,
		Name: owner.Name,
		UID:  owner.UID,
	}})
}

func injectFinalizer(object *metav1.ObjectMeta, finalizer string) {
	if len(object.Finalizers) == 0 {
		object.Finalizers = []string{
			finalizer,
		}
	} else {
		object.Finalizers = append(object.Finalizers, finalizer)
	}
}
