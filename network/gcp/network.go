package network

import (
	"bullion-pulumi/config"

	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataResult struct {
	Vpc        *compute.Network
	Subnetwork *compute.Subnetwork
	Router     *compute.Router
	Nat        *compute.RouterNat
	Firewall   []*compute.Firewall
	Resource   []pulumi.Resource
}

func Create(name string, ctx *pulumi.Context) (DataResult, error) {
	var result DataResult
	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)

	// create Vpc
	vpc, err := compute.NewNetwork(ctx, name, &compute.NetworkArgs{
		AutoCreateSubnetworks: pulumi.Bool(false),
	})
	if err != nil {
		return DataResult{}, err
	}

	// create subnet
	subnetwork, err := compute.NewSubnetwork(ctx, baseConfig.Subnetwork.Name, &compute.SubnetworkArgs{
		IpCidrRange: pulumi.String(baseConfig.Subnetwork.Cidr),
		Region:      pulumi.String(baseConfig.Subnetwork.Region),
		Network:     vpc.ID(),
	})
	if err != nil {
		return DataResult{}, err
	}

	// create router for NAT
	router, err := CreateRouter(ctx, "uji-router", subnetwork, vpc)
	if err != nil {
		return DataResult{}, err
	}

	// create NAT for accessing internet just use private IP
	nat, err := CreateNAT(ctx, "uji-nat", router)
	if err != nil {
		return DataResult{}, err
	}

	// Create ssh access
	sshAccess, err := SSHAccess(ctx, "ssh-access", vpc.Name, "0.0.0.0/0")
	if err != nil {
		return DataResult{}, err

	}

	// create mongo access
	mongoAccess, err := MongoAccess(ctx, "mongo-access", vpc.Name, "0.0.0.0/0")
	if err != nil {
		return DataResult{}, err

	}

	result.Vpc = vpc
	result.Subnetwork = subnetwork
	result.Router = router
	result.Nat = nat
	result.Firewall = append(result.Firewall, mongoAccess, sshAccess)
	result.Resource = append(result.Resource, vpc, subnetwork, router, nat, mongoAccess, sshAccess)

	return result, nil

}
