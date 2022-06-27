/*
Copyright 2022 hef.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"k8s.io/utils/env"

	cacctl "github.com/hef/cacctl/client"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	incubationv1beta1 "github.com/hef/cac-operator/api/v1beta1"
)

// ServerReconciler reconciles a Server object
type ServerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=incubation.hef.github.io,resources=servers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=incubation.hef.github.io,resources=servers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=incubation.hef.github.io,resources=servers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Server object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.1/pkg/reconcile
func (r *ServerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	x := incubationv1beta1.Server{}
	err := r.Get(ctx, req.NamespacedName, &x)
	if err != nil {
		logger.Error(err, "unable to fetch Server")
		return ctrl.Result{}, err
	}

	username := env.GetString("CAC_USERNAME", "")
	password := env.GetString("CAC_PASSWORD", "")

	client, err := cacctl.New(
		cacctl.WithUsernameAndPassword(username, password),
		cacctl.WithUserAgent("cac-operator/dev"),
	)
	if err != nil {
		logger.Error(err, "error creating client")
		return ctrl.Result{}, err
	}

	servers, err := client.ListWithFilter(ctx, req.Name, 200, cacctl.All)
	if err != nil {
		logger.Error(err, "error listing servers")
		return ctrl.Result{}, err
	}

	for _, server := range servers.Servers {
		if server.ServerName == req.Name {
			x.Status.Ip = server.IpAddress.String()
			x.Status.Id = server.ServerId
			x.Name = server.ServerName
		}
	}

	err = r.Status().Update(ctx, &x)
	if err != nil {
		logger.Error(err, "error updating server status")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ServerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&incubationv1beta1.Server{}).
		Complete(r)
}
