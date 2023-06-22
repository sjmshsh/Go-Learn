package main

import (
	"context"
	"fmt"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig *string
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	gvr := schema.GroupVersionResource{
		Version:  "v1",
		Resource: "pods",
	}
	unstructObj, err := dynamicClient.Resource(gvr).Namespace("gitlab").List(context.Background(), v1.ListOptions{Limit: 500})
	if err != nil {
		panic(err)
	}
	podList := &v12.PodList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructObj.UnstructuredContent(), podList)
	if err != nil {
		panic(err)
	}
	for _, d := range podList.Items {
		fmt.Println(d)
	}
}
