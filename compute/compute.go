package compute

import "github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"

type Compute interface {
	Create(params ...interface{}) (*compute.Instance, error)
}
