package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Configuration struct {
	Vpc        vpc
	Subnetwork subnetwork
	Compute    compute
	Firewall   firewall
	Kubernetes kubernetes
}

type vpc struct {
	Name   string
	Region string
}
type subnetwork struct {
	Name   string
	Region string
	Cidr   string
}

type compute struct {
	Zone string
}

type firewall struct {
	Range1 string
}

type kubernetes struct {
	Ip4cidr string
}

func GetConfig(baseConfig *Configuration) {
	basePath, _ := os.Getwd()
	if _, err := toml.DecodeFile(basePath+"/config.toml", &baseConfig); err != nil {
		fmt.Println(err)
	}
}
