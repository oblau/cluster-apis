package main 

import (
        "fmt"
        "sigs.k8s.io/controller-runtime/pkg/client/config"
        "sigs.k8s.io/controller-runtime/pkg/client"
        "k8s.io/klog"
)

func main() {
        client, _ := connect()
        fmt.Println(client)
}

func connect() (client.Client, error) {
     fmt.Println("we are in connect function")
     // by default it will look for KUBECONFIG env variable or check 
     // $HOME/.kube/config 
     cfg, err := config.GetConfig()
     if err != nil {
         fmt.Println("something bad happened, unable to connect to cluster")
         klog.Exit(err.Error())
     }
     c, err := client.New(cfg, client.Options{})
     if err != nil {
        klog.Exit(err.Error())
     }
     return c, nil 
}
