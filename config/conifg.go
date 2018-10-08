package config

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	Driver string
	Host string
	User string
	Password string
	Dbname string
	SslMode string
}