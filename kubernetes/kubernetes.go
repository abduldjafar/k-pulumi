package kubernetes

import (
	"bullion-pulumi/config"

	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/container"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NodePoolDetails struct {
	NodePoolName string
	ClusterName  pulumi.StringOutput
	Location     string
	MachineType  string
	NodeCount    int
}

type KubernetesResults struct {
	NodePools []*container.NodePool
}

func Create(ctx *pulumi.Context, kubename string, network pulumi.StringInput) (KubernetesResults, error) {
	var kubernetesResults KubernetesResults
	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)

	primary, err := container.NewCluster(ctx, kubename, &container.ClusterArgs{
		Location:              pulumi.String(baseConfig.Compute.Zone),
		RemoveDefaultNodePool: pulumi.Bool(true),
		InitialNodeCount:      pulumi.Int(1),
		Network:               network,
	})
	if err != nil {
		return kubernetesResults, err
	}

	for _, nodePool := range []NodePoolDetails{
		NodePoolDetails{
			NodePoolName: "node-1",
			ClusterName:  primary.Name,
			Location:     baseConfig.Compute.Zone,
			MachineType:  "e2-small",
			NodeCount:    1,
		},
	} {
		result, err := CreateNodePool(ctx, nodePool.NodePoolName, nodePool.ClusterName, nodePool.Location, nodePool.MachineType, nodePool.NodeCount)
		if err != nil {
			return kubernetesResults, err
		}

		kubernetesResults.NodePools = append(kubernetesResults.NodePools, result)
	}

	return kubernetesResults, nil

}
