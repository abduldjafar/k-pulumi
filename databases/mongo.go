package databases

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MongoDB interface {
	MongosCluster(params ...interface{}) ([]pulumi.Resource, error)
}
