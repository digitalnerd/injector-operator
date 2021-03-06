/*
Copyright 2022.

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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"github.com/digitalnerd/injector-operator/api/v1alpha1"
	injectorv1alpha1 "github.com/digitalnerd/injector-operator/api/v1alpha1"
)

// InjectorReconciler reconciles a Injector object
type InjectorReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=injector.dev.digitalnerd,resources=injectors,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=injector.dev.digitalnerd,resources=injectors/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=injector.dev.digitalnerd,resources=injectors/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Injector object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *InjectorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	var injector *v1alpha1.Injector
	r.Client.Get(ctx, req.NamespacedName, injector)
	if err := r.Get(ctx, req.NamespacedName, injector); err != nil {
		log.Log.Error(err, "enable to fetch a CronJob")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	myFanalizerName := "digitalnerd.injector.dev/finalizer"
	// var injectorList v1alpha1.InjectorList
	// r.Client.List(ctx, &injectorList)
	// for i,v := range injectorList.Items {
	// 	// ...
	// }

	
}

// SetupWithManager sets up the controller with the Manager.
func (r *InjectorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&injectorv1alpha1.Injector{}).
		Watches(
			&source.Kind{Type: &corev1.Pod{}},
			handler.EnqueueRequestsFromMapFunc(r.GetAll),
		).
		Complete(r)
}

func (r *InjectorReconciler) GetAll(o client.Client) []ctrl.Request {

}