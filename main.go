package main

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

func main() {
	// Kubernetes apiserver bağlantısını kur
	config, err := clientcmd.BuildConfigFromFlags("", "service.yaml")
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Tüm podları listele
	/*	pods, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		} */

	/* fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	for _, pod := range pods.Items {
		fmt.Printf("Name %s\n", pod.GetName())
		fmt.Printf("Status Condination %s\n", pod.Status.Phase)
	} */

	watchlist := cache.NewListWatchFromClient(
		clientset.CoreV1().RESTClient(),
		"pods",
		"",
		fields.Everything(),
	)

	_, controller := cache.NewInformer( // also take a look at NewSharedIndexInformer
		watchlist,
		&v1.Pod{},
		0, //Duration is int64
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				fmt.Printf("Yeni bir pod eklendi Pod Name: %s \n", obj.(*v1.Pod).Name)
				if obj.(*v1.Pod).Status.Phase == "Pending" {
					TeamsServer("Deneme Mesaj",
						"Yeni bir versiyon eklendi versiyon name : `"+obj.(*v1.Pod).Name+"`")
				}
			},
			DeleteFunc: func(obj interface{}) {
				fmt.Printf("Var olan bir pod silindi Pod Name: %s \n", obj.(*v1.Pod).Name)
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				fmt.Printf("Var olan bir pod güncellendi Pod Name: %s \n", newObj.(*v1.Pod).Name)
			},
		},
	)
	stop := make(chan struct{})
	defer close(stop)
	go controller.Run(stop)
	for {
		time.Sleep(time.Second)
	}

}
