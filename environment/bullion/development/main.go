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
		_, err = kubernetes.Create(ctx, "bullion-created-by-pulumi", result.Vpc.Name)
		if err != nil {
			return err

		}

		return nil
	})
}
