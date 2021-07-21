package compute

import (
	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Create(
	ctx *pulumi.Context,
	instanceName string,
	machineType string,
	network pulumi.IDOutput,
	subnetwork pulumi.IDOutput,
	zone string,
	os string,
	script string,
	size int,
	tags string,
	dependsOn ...[]pulumi.Resource,
) (
	*compute.Instance,
	error,
) {

	result, err := compute.NewInstance(ctx, instanceName, &compute.InstanceArgs{
		MachineType: pulumi.String(machineType),
		Zone:        pulumi.String(zone),
		Tags: pulumi.StringArray{
			pulumi.String(tags),
		},
		BootDisk: &compute.InstanceBootDiskArgs{
			InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
				Image: pulumi.String(os),
				Size:  pulumi.IntPtr(size),
			},
		},
		NetworkInterfaces: compute.InstanceNetworkInterfaceArray{
			&compute.InstanceNetworkInterfaceArgs{
				Network: network,
				AccessConfigs: compute.InstanceNetworkInterfaceAccessConfigArray{
					nil,
				},
				Subnetwork: subnetwork,
			},
		},
		MetadataStartupScript: pulumi.String(script),
		Metadata: pulumi.StringMap{
			"foo": pulumi.String("bar"),
		},
	}, pulumi.DependsOn(dependsOn[0]))
	if err != nil {
		return result, err
	}
	return result, nil

}
