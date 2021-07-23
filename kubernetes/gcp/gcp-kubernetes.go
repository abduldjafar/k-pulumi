package kubernetes

import (
	"bullion-pulumi/config"
	mykube "bullion-pulumi/kubernetes"

	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/container"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type gcpKubernetes struct{}
type KubernetesResults struct {
	NodePools []*container.NodePool
}

func (*gcpKubernetes) CreatePrivateGKE(params ...interface{}) error {
	ctx := params[0].(*pulumi.Context)
	kubename := params[1].(string)
	network := params[2].(pulumi.StringInput)
	subnetwork := params[3].(pulumi.StringInput)

	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)

	primary, err := container.NewCluster(ctx, kubename, &container.ClusterArgs{
		Location:              pulumi.String(baseConfig.Compute.Zone),
		InitialNodeCount:      pulumi.Int(1),
		RemoveDefaultNodePool: pulumi.Bool(true),
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
		return err
	}

	nodePoolsConfig, err := ReadConfig()
	if err != nil {
		return err
	}

	for _, config := range nodePoolsConfig {
		config.Config.Cluster = primary.Name
		config.Config.Location = pulumi.String(baseConfig.Compute.Zone)

		node, err := container.NewNodePool(ctx, config.NodePoolName, config.Config, pulumi.DependsOn([]pulumi.Resource{primary}))
		if err != nil {
			return err
		}
		ctx.Export("nodepool :", node)
	}

	return nil
}
func GCPKubernetes() mykube.Kubernetes {
	return &gcpKubernetes{}
}
