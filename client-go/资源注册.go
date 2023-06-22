package main

import (
	v12 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	v13 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func main() {
	// KnowType external
	coreGV := schema.GroupVersion{
		Group:   "",
		Version: "v1",
	}
	extensionsGV := schema.GroupVersion{
		Group:   "extensions",
		Version: "v1beta1",
	}

	coreInternalGV := schema.GroupVersion{
		Group:   "",
		Version: runtime.APIVersionInternal,
	}
	
	// UnversionedType
	Unversioned := schema.GroupVersion{
		Group:   "",
		Version: "v1",
	}

	schema := runtime.NewScheme()
	schema.AddKnownTypes(coreGV, &v1.Pod{})
	schema.AddKnownTypes(extensionsGV, &v12.DaemonSet{})
	schema.AddKnownTypes(coreInternalGV, &v1.Pod{})
	schema.AddUnversionedTypes(Unversioned, &v13.Status{})
}
