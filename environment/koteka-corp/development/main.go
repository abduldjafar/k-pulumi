package main

import (
	"k-pulumi/databases"
	"k-pulumi/kubernetes"
	"k-pulumi/network"

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

		// create mongo access
		mongoFirewall, err := network.MongoAccess(ctx, "mongo-access", result.Vpc.Name, "0.0.0.0/0")
		if err != nil {
			return err

		}

		// Create mongodb cluster
		_, err = databases.MongosCluster(ctx, result.Vpc.ID(), result.Subnetwork.ID(), []pulumi.Resource{result.Router, result.Nat, mongoFirewall})
		if err != nil {
			return err

		}

		// create Kubernetes
		nodePoolsDetails := []kubernetes.NodePoolDetails{
			kubernetes.NodePoolDetails{
				NodePoolName: "pool-1",
				Location:     "asia-southeast1",
				MachineType:  "e2-small",
				NodeCount:    1,
			},
		}

		_, err = kubernetes.CreatePrivateGKE(ctx, "development", result.Vpc.Name, result.Subnetwork.Name, nodePoolsDetails)
		if err != nil {
			return err

		}

		return nil
	})
}
