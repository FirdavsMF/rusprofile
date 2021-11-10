package config

import "github.com/kelseyhightower/envconfig"

func GetGrpcServerConfig() (GrpcServerConfig, error) {
	var cfg GrpcServerConfig
	err := envconfig.Process("rusprofile", &cfg.RusProfile)
	if err != nil {
		return GrpcServerConfig{}, err
	}

	err = envconfig.Process("grpc", &cfg.Grpc)
	if err != nil {
		return GrpcServerConfig{}, err
	}

	return cfg, nil
}

func GetHttpServerConfig() (HttpServerConfig, error) {
	var cfg HttpServerConfig
	err := envconfig.Process("http", &cfg.Http)
	if err != nil {
		return HttpServerConfig{}, err
	}

	err = envconfig.Process("grpc", &cfg.Grpc)
	if err != nil {
		return HttpServerConfig{}, err
	}

	return cfg, nil
}
