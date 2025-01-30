package config

type Config struct {
	GRPC CRPCConfig
	PG   PGConfig
}

type CRPCConfig struct {
	Port string `env:"GRPC_PORT"`
}

type PGConfig struct {
	DSN string `env:"POSTGRES_DSN"`
}
