package network

import (
	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateRouter(ctx *pulumi.Context, name string, subnet *compute.Subnetwork, net *compute.Network) (*compute.Router, error) {
	router, err := compute.NewRouter(ctx, name, &compute.RouterArgs{
		Region:  subnet.Region,
		Network: net.ID(),
		Bgp: &compute.RouterBgpArgs{
			Asn: pulumi.Int(64514),
		},
	})
	if err != nil {
		return &compute.Router{}, err
	}

	return router, nil
}
