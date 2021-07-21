package databases

import (
	computes "k-pulumi/compute"
	"k-pulumi/config"
	"log"

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

func MongosCluster(ctx *pulumi.Context, vpc pulumi.IDOutput, subnetwork pulumi.IDOutput, dependsOn ...[]pulumi.Resource) (Result, error) {
	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)

	var results Result

	for _, computeDetail := range []ComputeDetailes{
		ComputeDetailes{
			Name:        "mongo-master",
			MachineType: "e2-medium",
			VpcId:       vpc,
			Subnetwork:  subnetwork,
			Zone:        baseConfig.Compute.Zone,
			Os:          "debian-cloud/debian-10",
			MetadataStartupScript: `#!/bin/bash
			sudo apt install wget -y
			wget -qO - https://www.mongodb.org/static/pgp/server-5.0.asc | sudo apt-key add -
			echo "deb http://repo.mongodb.org/apt/debian buster/mongodb-org/5.0 main" | sudo tee /etc/apt/sources.list.d/mongodb-org-5.0.list
			sudo apt-get update
			sudo apt-get install -y mongodb-org
			sudo systemctl start mongod
			sudo systemctl enable mongod
			sudo systemctl stop mongod
			sudo rm /tmp/mongo*
			sudo nohup mongod --replSet "rs0" --bind_ip 0.0.0.0 --dbpath /var/lib/mongodb &`,
			Size: 20,
			Tags: "mongo-master",
		},
		ComputeDetailes{
			Name:        "mongo-slave-1",
			MachineType: "e2-small",
			VpcId:       vpc,
			Subnetwork:  subnetwork,
			Zone:        baseConfig.Compute.Zone,
			Os:          "debian-cloud/debian-10",
			MetadataStartupScript: `#!/bin/bash
			sudo apt install wget -y
			wget -qO - https://www.mongodb.org/static/pgp/server-5.0.asc | sudo apt-key add -
			echo "deb http://repo.mongodb.org/apt/debian buster/mongodb-org/5.0 main" | sudo tee /etc/apt/sources.list.d/mongodb-org-5.0.list
			sudo apt-get update
			sudo apt-get install -y mongodb-org
			sudo systemctl start mongod
			sudo systemctl enable mongod
			sudo systemctl stop mongod
			sudo rm /tmp/mongo*
			sudo nohup mongod --replSet "rs0" --bind_ip 0.0.0.0 --dbpath /var/lib/mongodb &`,
			Size: 30,
			Tags: "mongo-slave",
		},
		ComputeDetailes{
			Name:        "mongo-slave-2",
			MachineType: "e2-small",
			VpcId:       vpc,
			Subnetwork:  subnetwork,
			Zone:        baseConfig.Compute.Zone,
			Os:          "debian-cloud/debian-10",
			MetadataStartupScript: `#!/bin/bash
			sudo apt install wget -y
			wget -qO - https://www.mongodb.org/static/pgp/server-5.0.asc | sudo apt-key add -
			echo "deb http://repo.mongodb.org/apt/debian buster/mongodb-org/5.0 main" | sudo tee /etc/apt/sources.list.d/mongodb-org-5.0.list
			sudo apt-get update
			sudo apt-get install -y mongodb-org
			sudo systemctl start mongod
			sudo systemctl enable mongod
			sudo systemctl stop mongod
			sudo rm /tmp/mongo*
			sudo nohup mongod --replSet "rs0" --bind_ip 0.0.0.0 --dbpath /var/lib/mongodb &`,
			Size: 25,
			Tags: "mongo-slave",
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
			computeDetail.Size, computeDetail.Tags,
			dependsOn[0],
		)
		if err != nil {
			return results, err

		}
		log.Println(result.Hostname)
		results.Master = append(results.Master, result)
	}

	return results, nil
}

func MongoSingle(ctx *pulumi.Context, vpc pulumi.IDOutput, subnetwork pulumi.IDOutput, dependsOn ...[]pulumi.Resource) (*compute.Instance, error) {
	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)

	computeDetail := ComputeDetailes{
		Name:        "mongo-master",
		MachineType: "e2-medium",
		VpcId:       vpc,
		Subnetwork:  subnetwork,
		Zone:        baseConfig.Compute.Zone,
		Os:          "debian-cloud/debian-10",
		MetadataStartupScript: `#!/bin/bash
								sudo apt install wget -y
								wget -qO - https://www.mongodb.org/static/pgp/server-5.0.asc | sudo apt-key add -
								echo "deb http://repo.mongodb.org/apt/debian buster/mongodb-org/5.0 main" | sudo tee /etc/apt/sources.list.d/mongodb-org-5.0.list
								sudo apt-get update
								sudo apt-get install -y mongodb-org
								sudo systemctl start mongod
								sudo systemctl enable mongod`,
		Size: 20,
		Tags: "mongo-master",
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
		dependsOn[0],
	)

	if err != nil {
		return &compute.Instance{}, err

	}

	return result, nil
}
