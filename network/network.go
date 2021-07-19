package network

import (
	"bullion-pulumi/config"

	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataResult struct {
	Vpc        *compute.Network
	Subnetwork *compute.Subnetwork
}

func Create(name string, ctx *pulumi.Context) (DataResult, error) {
	var result DataResult
	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)

	custom_test, err := compute.NewNetwork(ctx, name, &compute.NetworkArgs{
		AutoCreateSubnetworks: pulumi.Bool(false),
	})
	if err != nil {
		return DataResult{}, err
	}

	custom_test_2, err := compute.NewSubnetwork(ctx, baseConfig.Subnetwork.Name, &compute.SubnetworkArgs{
		IpCidrRange: pulumi.String(baseConfig.Subnetwork.Cidr),
		Region:      pulumi.String(baseConfig.Subnetwork.Region),
		Network:     custom_test.ID(),
	})

	if err != nil {
		return DataResult{}, err
	}

	result.Vpc = custom_test
	result.Subnetwork = custom_test_2

	return result, nil

}
