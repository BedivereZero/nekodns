package apiserver

type Configuration struct {
	BindAddress string `json:"bindAddress"`
	SecurePort  uint16 `json:"securePort"`
}

type Server struct {
	cfg *Configuration
}

func New(cfg *Configuration) *Server {
	return &Server{cfg: cfg}
}

func NewDefaultConfiguration() *Configuration {
	return &Configuration{
		BindAddress: DefaultBindAddress,
		SecurePort:  DefaultSecurePort,
	}
}
