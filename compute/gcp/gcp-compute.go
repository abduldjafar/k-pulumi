package compute

import (
	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	computes "bullion-pulumi/compute"
)

type gcpCompute struct{}

func (*gcpCompute) Create(params ...interface{}) (*compute.Instance, error) {
	ctx := params[0].(*pulumi.Context)
	instanceName := params[1].(string)
	machineType := params[2].(string)
	network := params[3].(pulumi.IDOutput)
	subnetwork := params[4].(pulumi.IDOutput)
	zone := params[5].(string)
	os := params[6].(string)
	script := params[7].(string)
	size := params[8].(int)
	tags := params[9].(string)
	dependsOn := params[10].([]pulumi.Resource)

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
	}, pulumi.DependsOn(dependsOn))
	if err != nil {
		return result, err
	}
	return result, nil

}

func GCPCompute() computes.Compute {
	return &gcpCompute{}
}
