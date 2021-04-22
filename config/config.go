package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port                                     int    `envconfig:"PORT" default:"3000"`
	DatabaseName                             string `envconfig:"DATABASE_NAME" default:"db.db"`
	DisableForeignKeyConstraintWhenMigrating bool   `envconfig:"DATABASE_CONSTRAINT_WHEN_MIGRATING" default:"true"` //to make it work on sqlite
}

func Load() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
