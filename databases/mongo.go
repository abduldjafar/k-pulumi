package databases

import (
	computes "bullion-pulumi/compute"
	"bullion-pulumi/config"

	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Result struct {
	Master []*compute.Instance
}

type ComputeDetailes struct {
	Name                  string
	MachineType           string
	VpcId                 pulumi.IDOutput
	Subnetwork            pulumi.IDOutput
	Zone                  string
	Os                    string
	MetadataStartupScript string
	Size                  int
	Tags                  string
}

func MongosCluster(ctx *pulumi.Context, vpc pulumi.IDOutput, subnetwork pulumi.IDOutput) (Result, error) {
	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)

	var results Result

	for _, computeDetail := range []ComputeDetailes{
		ComputeDetailes{
			Name:                  "mongo-master",
			MachineType:           "e2-medium",
			VpcId:                 vpc,
			Subnetwork:            subnetwork,
			Zone:                  baseConfig.Compute.Zone,
			Os:                    "debian-cloud/debian-10",
			MetadataStartupScript: "gsutil cp gs://decentralize-config/mongodb-installation.sh . && bash mongodb-installation.sh && sudo systemctl start mongod && sudo systemctl enable mongod",
			Size:                  20,
			Tags:                  "mongo-master",
		},
		ComputeDetailes{
			Name:                  "mongo-slave-1",
			MachineType:           "e2-small",
			VpcId:                 vpc,
			Subnetwork:            subnetwork,
			Zone:                  baseConfig.Compute.Zone,
			Os:                    "debian-cloud/debian-10",
			MetadataStartupScript: "gsutil cp gs://decentralize-config/mongodb-installation.sh . && bash mongodb-installation.sh && sudo systemctl start mongod && sudo systemctl enable mongod",
			Size:                  30,
			Tags:                  "mongo-small",
		},
		ComputeDetailes{
			Name:                  "mongo-slave-2",
			MachineType:           "e2-small",
			VpcId:                 vpc,
			Subnetwork:            subnetwork,
			Zone:                  baseConfig.Compute.Zone,
			Os:                    "debian-cloud/debian-10",
			MetadataStartupScript: "gsutil cp gs://decentralize-config/mongodb-installation.sh . && bash mongodb-installation.sh && sudo systemctl start mongod && sudo systemctl enable mongod",
			Size:                  25,
			Tags:                  "mongo-small",
		},
	} {
		result, err := computes.Create(
			ctx,
			computeDetail.Name,
			computeDetail.MachineType,
			computeDetail.VpcId,
			computeDetail.Subnetwork,
			computeDetail.Zone,
			computeDetail.Os,
			computeDetail.MetadataStartupScript,
			computeDetail.Size, computeDetail.Tags)
		if err != nil {
			return results, err

		}

		results.Master = append(results.Master, result)
	}

	return results, nil
}

func MongoSingle(ctx *pulumi.Context, vpc pulumi.IDOutput, subnetwork pulumi.IDOutput) (*compute.Instance, error) {
	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)

	computeDetail := ComputeDetailes{
		Name:                  "mongo-master",
		MachineType:           "e2-medium",
		VpcId:                 vpc,
		Subnetwork:            subnetwork,
		Zone:                  baseConfig.Compute.Zone,
		Os:                    "debian-cloud/debian-10",
		MetadataStartupScript: "gsutil cp gs://decentralize-config/mongodb-installation.sh . && bash mongodb-installation.sh && sudo systemctl start mongod && sudo systemctl enable mongod",
		Size:                  20,
		Tags:                  "mongo-master",
	}

	result, err := computes.Create(
		ctx,
		computeDetail.Name,
		computeDetail.MachineType,
		computeDetail.VpcId,
		computeDetail.Subnetwork,
		computeDetail.Zone,
		computeDetail.Os,
		computeDetail.MetadataStartupScript,
		computeDetail.Size, computeDetail.Tags,
	)

	if err != nil {
		return &compute.Instance{}, err

	}

	return result, nil
}
