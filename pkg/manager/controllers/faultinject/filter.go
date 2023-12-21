/*
Copyright 2023 The KusionStack Authors.

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

package faultinject

import (
	"sigs.k8s.io/controller-runtime/pkg/event"

	"github.com/KusionStack/controller-mesh/pkg/apis/ctrlmesh"
	"github.com/KusionStack/controller-mesh/pkg/apis/ctrlmesh/utils/conv"
	ctrlmeshv1alpha1 "github.com/KusionStack/controller-mesh/pkg/apis/ctrlmesh/v1alpha1"
)

type FaultInjectionPredicate struct {
}

// Create returns true if the Create event should be processed
func (b *FaultInjectionPredicate) Create(event.CreateEvent) bool {
	return true
}

// Delete returns true if the Delete event should be processed
func (b *FaultInjectionPredicate) Delete(e event.DeleteEvent) bool {
	return true
}

// Update returns true if the Update event should be processed
func (b *FaultInjectionPredicate) Update(e event.UpdateEvent) bool {
	oldFault := e.ObjectOld.(*ctrlmeshv1alpha1.FaultInjection)
	newFault := e.ObjectNew.(*ctrlmeshv1alpha1.FaultInjection)
	if newFault.DeletionTimestamp != nil || len(oldFault.Finalizers) != len(newFault.Finalizers) {
		return true
	}
	if newFault.Labels != nil {
		_, ok := newFault.Labels[ctrlmesh.CtrlmeshFaultInjectionDisableKey]
		if ok {
			return true
		}
	}
	oldProtoFault := conv.ConvertFaultInjection(oldFault)
	newProtoFault := conv.ConvertFaultInjection(newFault)
	return oldProtoFault.ConfigHash != newProtoFault.ConfigHash
}

// Generic returns true if the Generic event should be processed
func (b *FaultInjectionPredicate) Generic(event.GenericEvent) bool {
	return true
}
