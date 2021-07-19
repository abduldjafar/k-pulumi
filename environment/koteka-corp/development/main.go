package main

import (
	"bullion-pulumi/databases"
	"bullion-pulumi/kubernetes"
	"bullion-pulumi/network"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a GCP network
		result, err := network.Create("uji-pulumi", ctx)
		if err != nil {
			return err

		}

		// Create ssh access
		_, err = network.SSHAccess(ctx, "ssh-access", result.Vpc.Name, "0.0.0.0/0")
		if err != nil {
			return err

		}

		// Create mongodb cluster
		_, err = databases.MongoSingle(ctx, result.Vpc.ID(), result.Subnetwork.ID())
		if err != nil {
			return err

		}

		// create Kubernetes
		nodePoolsDetails := []kubernetes.NodePoolDetails{
			kubernetes.NodePoolDetails{
				NodePoolName: "pool-1",
				Location:     "asia-southeast1",
				MachineType:  "e2-small",
				NodeCount:    3,
			},
		}

		_, err = kubernetes.CreateGKE(ctx, "development", result.Vpc.Name, result.Subnetwork.Name, nodePoolsDetails)
		if err != nil {
			return err

		}

		return nil
	})
}
