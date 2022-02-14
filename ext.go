package config

import (
	"encoding/json"
	"strings"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

func isJSON(ext string) bool {
	return strings.EqualFold(ext, ".json")
}

func loadJSON(content []byte, value interface{}) error {
	return json.Unmarshal(content, value)
}

func isYAML(ext string) bool {
	return strings.EqualFold(ext, ".yaml") || strings.EqualFold(ext, ".yml")
}

func loadYAML(content []byte, value interface{}) error {
	return yaml.Unmarshal(content, value)
}

func isTOML(ext string) bool {
	return strings.EqualFold(ext, ".toml")
}

func loadTOML(content []byte, value interface{}) error {
	return toml.Unmarshal(content, value)
}
