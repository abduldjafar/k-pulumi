package main

import (
	"bullion-pulumi/databases"
	databasesProvider "bullion-pulumi/databases/gcp/mongodb"
	kubernetes "bullion-pulumi/kubernetes"
	kubernetes_provider "bullion-pulumi/kubernetes/gcp"
	network "bullion-pulumi/network/gcp"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var (
	mongoDB databases.MongoDB     = databasesProvider.GcpMongoDB()
	myKube  kubernetes.Kubernetes = kubernetes_provider.GCPKubernetes()
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a GCP network
		result, err := network.Create("uji-pulumi", ctx)
		if err != nil {
			return err

		}

		// Create mongodb cluster
		_, err = mongoDB.MongosCluster(ctx, result.Vpc.ID(), result.Subnetwork.ID(), result.Resource)
		if err != nil {
			return err

		}

		// create Kubernetes
		err = myKube.CreatePrivateGKE(ctx, "development", result.Vpc.Name, result.Subnetwork.Name)
		if err != nil {
			return err

		}

		return nil
	})
}
