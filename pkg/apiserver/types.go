package apiserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Server struct {
	Engine *gin.Engine
	Client *clientv3.Client
}

func New(cfg Config) (*Server, error) {
	etcdClient, err := clientv3.NewFromURLs(cfg.EtcdEndpoints)
	if err != nil {
		return nil, fmt.Errorf("create etcd client fail: %w", err)
	}

	s := &Server{
		Engine: gin.Default(),
		Client: etcdClient,
	}

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
