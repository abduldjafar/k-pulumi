package kubernetes

import (
	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/container"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateNodePool(ctx *pulumi.Context, NodePoolName string, clusterName pulumi.StringOutput, location string, machineType string, nodeCount int) (*container.NodePool, error) {
	nodePool, err := container.NewNodePool(ctx, NodePoolName, &container.NodePoolArgs{
		Location:  pulumi.String(location),
		Cluster:   clusterName,
		NodeCount: pulumi.Int(nodeCount),
		NodeConfig: &container.NodePoolNodeConfigArgs{
			Preemptible: pulumi.Bool(true),
			MachineType: pulumi.String(machineType),
			OauthScopes: pulumi.StringArray{
				pulumi.String("https://www.googleapis.com/auth/cloud-platform"),
			},
		},
	})
	if err != nil {
		return &container.NodePool{}, err
	}

	return nodePool, nil
}
