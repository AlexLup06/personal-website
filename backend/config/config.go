package config

type Config struct {
	DevMode bool `env:"DEV_MODE, default=true"`
}
