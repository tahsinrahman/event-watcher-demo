/*
Copyright 2019 The Kubernetes Authors.

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

package main

import (
	"os"

	v1 "k8s.io/api/core/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{})
	if err != nil {
		os.Exit(1)
	}

	myReconciler := &eventReconciler{}

	if err := ctrl.NewControllerManagedBy(mgr).For(&v1.Event{}).Complete(myReconciler); err != nil {
		os.Exit(1)
	}
}

type eventReconciler struct {
}

func (r *eventReconciler) Reconcile(request ctrl.Request) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}
