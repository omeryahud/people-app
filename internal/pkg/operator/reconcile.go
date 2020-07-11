package operator

import (
	"context"

	"github.com/go-logr/logr"
	peoplev1alpha1 "github.com/omeryahud/people-app/api/v1alpha1"
	backendresources "github.com/omeryahud/people-app/internal/pkg/backend/resources"
	databaseresources "github.com/omeryahud/people-app/internal/pkg/database/resources"
	frontendresources "github.com/omeryahud/people-app/internal/pkg/frontend/resources"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func Reconcile(client client.Client,
	log logr.Logger,
	scheme *runtime.Scheme,
	req ctrl.Request) (ctrl.Result, error) {
	context := context.Background()
	_ = log.WithValues("peopleapp", req.NamespacedName)

	instance := &peoplev1alpha1.PeopleApp{}

	err := client.Get(context, req.NamespacedName, instance)

	if err != nil {
		log.Error(err, "could not fetch CR instance in reconcile loop")
		return ctrl.Result{}, err
	}

	log.Info("Reconciling PeopleApp", "namespace:", req.Namespace, "name:", req.Name)

	err = client.Create(context, frontendresources.NewDeployment(req.Namespace,
		instance.Spec.FrontendConfig.Image,
		instance.Spec.FrontendConfig.Replicas))

	if err != nil {
		log.Error(err, "failed frontend deployment creation")
		return ctrl.Result{}, err
	}

	client.Create(context, backendresources.NewDeployment(req.Namespace,
		instance.Spec.BackendConfig.Image,
		instance.Spec.BackendConfig.Replicas))

	if err != nil {
		log.Error(err, "failed backend deployment creation")
		return ctrl.Result{}, err
	}

	client.Create(context, databaseresources.NewDeployment(req.Namespace,
		instance.Spec.DatabaseConfig.Image,
		instance.Spec.DatabaseConfig.Replicas))

	if err != nil {
		log.Error(err, "failed database deployment creation")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}
