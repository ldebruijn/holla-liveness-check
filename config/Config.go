package config

import "larsdebruijn.nl/holla/target"

type Config struct {
	Groups []target.Group `yaml:"groups"`
}
