// Package config to env
package config

import "github.com/caarlos0/env/v6"

// RedisConfig struct to redis config env
type RedisConfig struct {
	Addr     string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	Password string `env:"REDIS_PASSWORD" envDefault:""`
	DB       int    `env:"REDIS_DB" envDefault:"0"`
}

// NewRedis contract redis config
func NewRedis() (*RedisConfig, error) {
	cfg := &RedisConfig{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
