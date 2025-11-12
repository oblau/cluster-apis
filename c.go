package main

import (
	"context"
	"fmt"

	"github.com/openshift/cluster-node-tuning-operator/test/e2e/performanceprofile/functests/utils/log"

	"github.com/openshift/cluster-node-tuning-operator/test/e2e/performanceprofile/functests/utils/nodes"
)

func main() {
	log.SetLevel(log.LevelError)

	ctx := context.Background()
	workerRTNodes, err := nodes.GetByRole("worker")
	if err != nil {
		fmt.Println("cannot get nodes")
	}
	for i, worker := range workerRTNodes {
		fmt.Printf("worker #%v - %s\n", i, worker.Name)

		cpuSet, err := nodes.GetOnlineCPUsSet(ctx, &worker)
		if err != nil {
			fmt.Printf("cannot get online cpuset: %v\n", err)
		} else {
			fmt.Printf("online cpu set for this node is: %s\n", cpuSet.String())
		}

		err = nodes.HasPreemptRTKernel(ctx, &worker)
		if err != nil {
			fmt.Println("does not have RT kernel on this node")
		} else {
			fmt.Println("has RT kernel on this node")
		}

	}

	numaNodes, err := nodes.GetNumaNodes(ctx, &(workerRTNodes[0]))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("NUMA topology for first worker node:")
	for nodeID, cpus := range numaNodes {
		fmt.Printf("  NUMA Node %d: CPUs %v\n", nodeID, cpus)
	}
}
