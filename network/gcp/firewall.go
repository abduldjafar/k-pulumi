package network

import (
	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func SSHAccess(ctx *pulumi.Context, firewallName string, networkName pulumi.StringInput, sourceRanges string) (*compute.Firewall, error) {

	result, err := compute.NewFirewall(ctx, firewallName, &compute.FirewallArgs{
		Name:    pulumi.String(firewallName),
		Network: networkName,
		Allows: compute.FirewallAllowArray{
			&compute.FirewallAllowArgs{
				Protocol: pulumi.String("tcp"),
				Ports: pulumi.StringArray{
					pulumi.String("22"),
				},
			},
		},
		SourceRanges: pulumi.StringArray{
			pulumi.String(sourceRanges),
		},
	})
	if err != nil {
		return &compute.Firewall{}, err
	}

	return result, nil
}

func MongoAccess(ctx *pulumi.Context, firewallName string, networkName pulumi.StringInput, sourceRanges string) (*compute.Firewall, error) {

	result, err := compute.NewFirewall(ctx, firewallName, &compute.FirewallArgs{
		Name:    pulumi.String(firewallName),
		Network: networkName,
		Allows: compute.FirewallAllowArray{
			&compute.FirewallAllowArgs{
				Protocol: pulumi.String("tcp"),
				Ports: pulumi.StringArray{
					pulumi.String("27017"),
				},
			},
		},
		SourceRanges: pulumi.StringArray{
			pulumi.String(sourceRanges),
		},
	})
	if err != nil {
		return &compute.Firewall{}, err
	}

	return result, nil
}
