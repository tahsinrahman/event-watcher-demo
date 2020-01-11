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
	"context"
	"fmt"
	"os"

	"github.com/the-redback/go-oneliners"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func main() {
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	myReconciler := &eventReconciler{
		client: mgr.GetClient(),
	}

	if err := ctrl.NewControllerManagedBy(mgr).For(&v1.Event{}).Complete(myReconciler); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type eventReconciler struct {
	client client.Client
}

func (r *eventReconciler) Reconcile(request ctrl.Request) (ctrl.Result, error) {
	fmt.Println(request.Name, request.Namespace)

	event := &v1.Event{}
	if err := r.client.Get(context.Background(), request.NamespacedName, event); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	oneliners.PrettyJson(event)

	return ctrl.Result{}, nil
}
