package conf

type CloudCore struct {
	HTTP string
	Grpc string
}

type MicroService struct {
	CloudCore CloudCore
	Etcd      Etcd
}
type Etcd struct {
	Addr []string
}
