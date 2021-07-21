package network

import (
	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateNAT(ctx *pulumi.Context, name string, router *compute.Router) (*compute.RouterNat, error) {
	nat, err := compute.NewRouterNat(ctx, name, &compute.RouterNatArgs{
		Router:                        router.Name,
		Region:                        router.Region,
		NatIpAllocateOption:           pulumi.String("AUTO_ONLY"),
		SourceSubnetworkIpRangesToNat: pulumi.String("ALL_SUBNETWORKS_ALL_IP_RANGES"),
	})
	if err != nil {
		return &compute.RouterNat{}, err
	}

	return nat, nil
}
