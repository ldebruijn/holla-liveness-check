package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"larsdebruijn.nl/holla/target"
	"log"
	"os"
)

type YamlConfigurationService struct{}

func (c *YamlConfigurationService) Load(configuration chan<- []target.Group) {
	log.Println("Loading configuration")
	body := c.readFileContents(c.getFilePath())
	parsed := c.parseFileContents(body)

	configuration <- parsed.Groups
}

func (c *YamlConfigurationService) getFilePath() string {
	path := os.Getenv("HOLA_CONFIG_PATH")
	if path != "" {
		return path
	}

	return "./configs/config.yml"
}

func (c *YamlConfigurationService) readFileContents(path string) []byte {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Cannot read config file: %v", err)
	}

	log.Println(string(body))
	return body
}

func (c *YamlConfigurationService) parseFileContents(body []byte) Config {
	var parsed Config

	err := yaml.Unmarshal(body, &parsed)
	if err != nil {
		log.Fatalf("Unable to parse config file: %v", err)
	}
	return parsed
}
