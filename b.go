package main


import (
        "fmt"
        corev1 "k8s.io/api/core/v1"
        "github.com/openshift/cluster-node-tuning-operator/test/e2e/performanceprofile/functests/utils/nodes"
	"github.com/openshift/cluster-node-tuning-operator/test/e2e/performanceprofile/functests/utils/profiles"
)

func main() {
        var workerRTNodes []corev1.Node
        var err error
        workerRTNodes, err = nodes.GetByRole("worker-cnf")
        if err != nil {
                fmt.Println("Unable to get nodes")
        } else {
                for _, node := range workerRTNodes {
                        fmt.Println("nodes are: ", node.Name)
                }
        }
	workerLabel := make(map[string]string)
	workerLabel["node-role.kubernetes.io/worker-cnf"] = ""
	fmt.Println(workerLabel)
	performanceProfile, err := profiles.GetByNodeLabels(workerLabel)
	fmt.Println(*performanceProfile.Spec.CPU.Reserved)
}
