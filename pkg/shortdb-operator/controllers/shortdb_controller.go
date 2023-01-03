/*
Copyright 2022 Viktor Login.

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
	"strconv"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	shortdbv1alpha1 "github.com/shortlink-org/shortlink/pkg/shortdb-operator/api/v1alpha1"
)

// ShortDBReconciler reconciles a ShortDB object
type ShortDBReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=shortdb.shortdb.shortlink,resources=shortdbs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=shortdb.shortdb.shortlink,resources=shortdbs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=shortdb.shortdb.shortlink,resources=shortdbs/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps/v1,resources=statefulsets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps/v1,resources=pods,verbs=get;list
//+kubebuilder:rbac:groups=v1,resources=configmaps,verbs=get;list;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ShortDB object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.1/pkg/reconcile
func (r *ShortDBReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// get the ShortDB object
	cluster := &shortdbv1alpha1.ShortDB{}
	err := r.Get(ctx, req.NamespacedName, cluster)
	if err != nil {
		return ctrl.Result{}, err
	}

	// validate the object
	var deployments int32
	deployments = int32(cluster.Spec.Deployments)
	var maxMemory string
	maxMemory = strconv.Itoa(*cluster.Spec.MaxMemory)
	var maxCpu string
	maxCpu = strconv.Itoa(*cluster.Spec.MaxCPU)

	// create a new statefulset
	sts := appsv1.StatefulSet{ObjectMeta: ctrl.ObjectMeta{Namespace: req.NamespacedName.Namespace, Name: req.NamespacedName.Name}, Spec: appsv1.StatefulSetSpec{
		Replicas: &deployments,
		Selector: &v1.LabelSelector{
			MatchLabels: map[string]string{"label": req.Name},
		},
		Template: corev1.PodTemplateSpec{
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:    "shortdb",
						Image:   "batazor/shortdb",
						Args:    nil,
						EnvFrom: nil,
						Env: []corev1.EnvVar{
							{
								Name:  "MEM_GB",
								Value: maxMemory,
							},
						},
						Resources: corev1.ResourceRequirements{
							Limits: corev1.ResourceList{
								corev1.ResourceCPU: resource.MustParse(maxCpu),
							},
						},
					},
				},
			},
		},
		UpdateStrategy: appsv1.StatefulSetUpdateStrategy{},
	}}

	err = r.Create(ctx, &sts)
	if err != nil {
		return ctrl.Result{
			Requeue: false,
		}, err
	}

	return ctrl.Result{
		Requeue: true,
	}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ShortDBReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&shortdbv1alpha1.ShortDB{}).
		Owns(&appsv1.StatefulSet{}).
		Owns(&corev1.ConfigMap{}).
		Complete(r)
}
