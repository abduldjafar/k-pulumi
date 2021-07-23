package kubernetes

type Kubernetes interface {
	CreatePrivateGKE(params ...interface{}) error
}
