// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// This file was autogenerated by defaulter-gen. Do not edit it manually!

package v1beta1

import (
	v1beta1 "k8s.io/api/storage/v1beta1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&v1beta1.StorageClass{}, func(obj interface{}) { SetObjectDefaults_StorageClass(obj.(*v1beta1.StorageClass)) })
	scheme.AddTypeDefaultingFunc(&v1beta1.StorageClassList{}, func(obj interface{}) { SetObjectDefaults_StorageClassList(obj.(*v1beta1.StorageClassList)) })
	return nil
}

func SetObjectDefaults_StorageClass(in *v1beta1.StorageClass) {
	SetDefaults_StorageClass(in)
}

func SetObjectDefaults_StorageClassList(in *v1beta1.StorageClassList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_StorageClass(a)
	}
}
