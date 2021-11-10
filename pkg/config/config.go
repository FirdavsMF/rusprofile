package config

type HttpServerConfig struct {
	Http Http
	Grpc Grpc
}

type GrpcServerConfig struct {
	RusProfile RusProfile
	Grpc       Grpc
}

type RusProfile struct {
	Host string `required:"true"`
}

type Http struct {
	Port string `required:"true"`
}

type Grpc struct {
	Port string `required:"true"`
}
