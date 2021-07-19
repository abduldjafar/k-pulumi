package main

import (
	"bullion-pulumi/databases"
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
		results, err := databases.Mongo(ctx, result.Vpc.ID(), result.Subnetwork.ID())
		if err != nil {
			return err

		}

		// Export the ids instances name of the bucket
		for _, data := range results.Master {
			ctx.Export("instances id created :", data.InstanceId)
		}

		// Export the vpc id
		ctx.Export("vpc-id", result.Vpc.ID())
		return nil
	})
}
