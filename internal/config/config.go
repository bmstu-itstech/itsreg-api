package config

import "time"

type Config struct {
	Env string `yaml:"env"    env:"ENV"         env-default:"local"`

	Server     ServerConfig `yaml:"server" env-required:"true"`
	BotsConfig BotsConfig   `yaml:"bots"   env-required:"true"`
	TgConfig   TgConfig     `yaml:"tg"     env-required:"true"`
}

type ServerConfig struct {
	Host        string        `yaml:"host"         env:"SRV_HOST"    env-default:"localhost"`
	Port        int           `yaml:"port"         env:"SRC_PORT"    env-default:"8000"`
	Timeout     time.Duration `yaml:"timeout"      env:"SRC_TIMEOUT" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"SRV_IDLE"    env-default:"10s"`
}

type AuthConfig struct {
	Host string `yaml:"host" env:"AUTH_HOST" env-default:"localhost"`
	Port int    `yaml:"port" env:"AUTH_PORT" env-default:"40001"`
}

type BotsConfig struct {
	Host string `yaml:"host" env:"BOTS_HOST" env-default:"localhost"`
	Port int    `yaml:"port" env:"BOTS_PORT" env-default:"40002"`
}

type TgConfig struct {
	Host string `yaml:"host" env:"TG_HOST" env-default:"localhost"`
	Port int    `yaml:"port" env:"TG_PORT" env-default:"40003"`
}
