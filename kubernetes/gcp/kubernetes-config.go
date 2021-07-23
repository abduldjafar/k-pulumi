package kubernetes

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/container"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NodePoolDetails struct {
	NodePoolName string
	Config       *container.NodePoolArgs
}

type KubernetesConfig struct {
	Kubernetes []struct {
		Poolname    string `json:"poolname"`
		Nodecount   int    `json:"nodecount"`
		Location    string `json:"location"`
		Machinetype string `json:"machinetype,omitempty"`
	} `json:"kubernetes"`
}

func ReadConfig() ([]NodePoolDetails, error) {
	var nodePoolDetails []NodePoolDetails
	var datas KubernetesConfig

	file, _ := ioutil.ReadFile("config.json")

	if err := json.Unmarshal([]byte(file), &datas); err != nil {

		return nodePoolDetails, err
	}
	for _, data := range datas.Kubernetes {
		nodePool := NodePoolDetails{
			NodePoolName: data.Poolname,
			Config: &container.NodePoolArgs{
				NodeCount: pulumi.Int(data.Nodecount),
				NodeConfig: &container.NodePoolNodeConfigArgs{
					Preemptible: pulumi.Bool(true),
					MachineType: pulumi.String(data.Machinetype),
				},
				Location: pulumi.String(data.Location),
			},
		}
		nodePoolDetails = append(nodePoolDetails, nodePool)
	}

	return nodePoolDetails, nil
}
