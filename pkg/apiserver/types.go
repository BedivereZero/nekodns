package apiserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Server struct {
	Router *gin.Engine
	Client *clientv3.Client
}

func New(cfg Config) (*Server, error) {
	etcdClient, err := clientv3.NewFromURLs(cfg.EtcdEndpoints)
	if err != nil {
		return nil, fmt.Errorf("create etcd client fail: %w", err)
	}

	s := &Server{
		Router: gin.Default(),
		Client: etcdClient,
	}

	RouterGroupV1(s)

	return s, nil
}

func Default() (*Server, error) {
	return New(Config{
		BindAddress: DefaultBindAddress,
		SecurePort:  DefaultSecurePort,
		EtcdEndpoints: []string{
			DefaultEtcdEndpoint,
		},
	})
}
