package config

type DB struct {
	Host       string
	Port       int
	Username   string
	Password   string
	DriverName string
	Name       string
}

type Config struct {
	AppGRPCPort int
	AppHTTPPort int
	AppEnv      string
	Debug       bool
	DB          *DB
}
