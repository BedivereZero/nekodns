package apiserver

type Config struct {
	BindAddress   string   `json:"bindAddress"`
	SecurePort    uint16   `json:"securePort"`
	EtcdEndpoints []string `json:"etcdEndpoints"`
}
