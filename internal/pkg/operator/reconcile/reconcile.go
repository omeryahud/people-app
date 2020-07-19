package reconcile

import (
	"context"
	"os"

	"github.com/go-logr/logr"
	peoplev1alpha1 "github.com/omeryahud/people-app/api/v1alpha1"
	"github.com/omeryahud/people-app/internal/pkg/operator"
	backendresources "github.com/omeryahud/people-app/internal/pkg/operator/resources/backend"
	databaseresources "github.com/omeryahud/people-app/internal/pkg/operator/resources/database"
	frontendresources "github.com/omeryahud/people-app/internal/pkg/operator/resources/frontend"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func Reconcile(client client.Client,
	log logr.Logger,
	scheme *runtime.Scheme,
	req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	_ = log.WithValues("peopleapp", req.NamespacedName)

	instance := &peoplev1alpha1.PeopleApp{}

	err := client.Get(ctx, req.NamespacedName, instance)

	if err != nil {
		log.Error(err, "could not fetch CR instance in reconcile loop")
		return ctrl.Result{}, err
	}

	log.Info("Reconciling PeopleApp", "namespace:", req.Namespace, "name:", req.Name)

	if err = reconcileFrontend(instance, client, log); err != nil {
		return ctrl.Result{}, err
	}

	if err = reconcileBackend(instance, client, log); err != nil {
		return ctrl.Result{}, err
	}

	if err = reconcileDatabase(instance, client, log); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func reconcileFrontend(instance *peoplev1alpha1.PeopleApp,
	client client.Client,
	log logr.Logger) error {
	frontendImage, _ := os.LookupEnv(operator.FrontendImageKey)
	replicas := func() int32 {
		if instance.Spec.FrontendSpec.Replicas != nil {
			return *instance.Spec.FrontendSpec.Replicas
		}

		return operator.DefaultFrontendReplicas
	}()

	ctx := context.TODO()

	deployment := frontendresources.NewDeployment(instance.Namespace,
		frontendImage,
		instance.Spec.FrontendSpec.HttpPort,
		replicas)
	injectOwnerReference(instance, &deployment.ObjectMeta)

	err := createOrPatchDeployment(client, ctx, deployment)
	if err != nil {
		return err
	}

	service := frontendresources.NewService(instance.Namespace,
		instance.Spec.FrontendSpec.HttpPort)
	injectOwnerReference(instance, &service.ObjectMeta)

	err = createOrPatchService(client, ctx, service)
	if err != nil {
		return err
	}

	return nil
}

func reconcileBackend(instance *peoplev1alpha1.PeopleApp,
	client client.Client,
	log logr.Logger) error {
	backendImage, _ := os.LookupEnv(operator.BackendImageKey)
	replicas := func() int32 {
		if instance.Spec.BackendSpec.Replicas != nil {
			return *instance.Spec.BackendSpec.Replicas
		}

		return operator.DefaultBackendReplicas
	}()

	ctx := context.TODO()

	deployment := backendresources.NewDeployment(instance.Namespace,
		backendImage,
		instance.Spec.BackendSpec.HttpPort,
		replicas)
	injectOwnerReference(instance, &deployment.ObjectMeta)

	err := createOrPatchDeployment(client, ctx, deployment)
	if err != nil {
		return err
	}

	service := backendresources.NewService(instance.Namespace,
		instance.Spec.BackendSpec.HttpPort)
	injectOwnerReference(instance, &service.ObjectMeta)

	err = createOrPatchService(client, ctx, service)
	if err != nil {
		return err
	}

	return nil
}

func reconcileDatabase(instance *peoplev1alpha1.PeopleApp,
	client client.Client,
	log logr.Logger) error {
	databaseImage, _ := os.LookupEnv(operator.DatabaseImageKey)
	replicas := func() int32 {
		if instance.Spec.DatabaseSpec.Replicas != nil {
			return *instance.Spec.DatabaseSpec.Replicas
		}

		return operator.DefaultDatabaseReplicas
	}()

	ctx := context.TODO()

	deployment := databaseresources.NewDeployment(instance.Namespace,
		databaseImage,
		instance.Spec.DatabaseSpec.HttpPort,
		replicas)
	injectOwnerReference(instance, &deployment.ObjectMeta)

	err := createOrPatchDeployment(client, ctx, deployment)
	if err != nil {
		return err
	}

	service := databaseresources.NewService(instance.Namespace,
		instance.Spec.DatabaseSpec.HttpPort)
	injectOwnerReference(instance, &service.ObjectMeta)

	err = createOrPatchService(client, ctx, service)
	if err != nil {
		return err
	}

	return nil
}
