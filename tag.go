package config

import (
	"fmt"
	"strings"
)

type tagOptions struct {
	Required     bool
	HasDefault   bool
	DefaultValue string
}

func parseTag(tag string) (*tagOptions, error) {
	options := tagOptions{}
	if len(tag) == 0 {
		return &options, nil
	}
	for _, s := range strings.Split(tag, ";") {
		switch {
		case s == "required":
			options.Required = true
		case strings.HasPrefix(s, "default="):
			defaultValue := s[len("default="):]
			if defaultValue != "-" {
				options.HasDefault = true
				options.DefaultValue = defaultValue
			}
		default:
			return nil, fmt.Errorf("unsupported option: \"%v\"", s)
		}
	}
	return &options, nil
}
