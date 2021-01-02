package config

import (
	"larsdebruijn.nl/holla/target"
)

type ConfigurationService interface {
	Load(configuration chan<- []target.Group)
}
