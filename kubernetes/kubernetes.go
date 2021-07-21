package kubernetes

import (
	"k-pulumi/config"

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

func CreateGKE(ctx *pulumi.Context, kubename string, network pulumi.StringInput, subnetwork pulumi.StringInput, nodePooldetails []NodePoolDetails) (KubernetesResults, error) {
	var kubernetesResults KubernetesResults
	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)

	primary, err := container.NewCluster(ctx, kubename, &container.ClusterArgs{
		Location:              pulumi.String(baseConfig.Compute.Zone),
		RemoveDefaultNodePool: pulumi.Bool(true),
		InitialNodeCount:      pulumi.Int(1),
		Network:               network,
		Subnetwork:            subnetwork,
	})
	if err != nil {
		return kubernetesResults, err
	}

	for _, nodePool := range nodePooldetails {
		nodePool.ClusterName = primary.Name

		result, err := CreateNodePool(ctx, nodePool.NodePoolName, nodePool.ClusterName, nodePool.Location, nodePool.MachineType, nodePool.NodeCount, []pulumi.Resource{primary})
		if err != nil {
			return kubernetesResults, err
		}

		kubernetesResults.NodePools = append(kubernetesResults.NodePools, result)
	}

	return kubernetesResults, nil

}

func CreatePrivateGKE(ctx *pulumi.Context, kubename string, network pulumi.StringInput, subnetwork pulumi.StringInput, nodePooldetails []NodePoolDetails) (KubernetesResults, error) {
	var kubernetesResults KubernetesResults
	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)

	primary, err := container.NewCluster(ctx, kubename, &container.ClusterArgs{
		Location:              pulumi.String(baseConfig.Compute.Zone),
		RemoveDefaultNodePool: pulumi.Bool(true),
		InitialNodeCount:      pulumi.Int(1),
		Network:               network,
		Subnetwork:            subnetwork,
		PrivateClusterConfig: &container.ClusterPrivateClusterConfigArgs{
			EnablePrivateEndpoint: pulumi.Bool(true),
			EnablePrivateNodes:    pulumi.Bool(true),
			MasterIpv4CidrBlock:   pulumi.String("172.16.2.0/28"),
			MasterGlobalAccessConfig: &container.ClusterPrivateClusterConfigMasterGlobalAccessConfigArgs{
				Enabled: pulumi.Bool(true),
			},
		},
		IpAllocationPolicy: &container.ClusterIpAllocationPolicyArgs{},
		MasterAuthorizedNetworksConfig: &container.ClusterMasterAuthorizedNetworksConfigArgs{
			CidrBlocks: &container.ClusterMasterAuthorizedNetworksConfigCidrBlockArray{
				&container.ClusterMasterAuthorizedNetworksConfigCidrBlockArgs{
					CidrBlock: pulumi.String("10.2.0.0/16"),
				},
			},
		},
	})
	if err != nil {
		return kubernetesResults, err
	}

	for _, nodePool := range nodePooldetails {
		nodePool.ClusterName = primary.Name

		result, err := CreateNodePool(
			ctx,
			nodePool.NodePoolName,
			nodePool.ClusterName,
			nodePool.Location,
			nodePool.MachineType,
			nodePool.NodeCount,
			[]pulumi.Resource{primary},
		)
		if err != nil {
			return kubernetesResults, err
		}

		kubernetesResults.NodePools = append(kubernetesResults.NodePools, result)
	}

	return kubernetesResults, nil

}
