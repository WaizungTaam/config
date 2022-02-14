package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
)

// Load loads config from a file
func Load(file string, config interface{}) error {
	if err := readConfig(file, config); err != nil {
		return err
	}
	if err := process(config); err != nil {
		return err
	}
	return nil
}

func readConfig(file string, config interface{}) error {
	value := reflect.Indirect(reflect.ValueOf(config))
	if !value.CanAddr() {
		return fmt.Errorf("config not addressable")
	}

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	ext := filepath.Ext(file)
	switch {
	case isJSON(ext):
		return loadJSON(content, config)
	case isYAML(ext):
		return loadYAML(content, config)
	case isTOML(ext):
		return loadTOML(content, config)
	default:
		return fmt.Errorf("unsupported file type")
	}
}

func process(config interface{}) error {
	configValue := reflect.Indirect(reflect.ValueOf(config))
	if configValue.Kind() != reflect.Struct {
		return fmt.Errorf("struct required")
	}
	configType := configValue.Type()
	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		value := configValue.Field(i)

		tag := field.Tag.Get("config")
		options, err := parseTag(tag)
		if err != nil {
			return err
		}
		isZeroValue := isZero(value)

		if options.Required && isZeroValue {
			return fmt.Errorf("field \"%s\" required", field.Name)
		}

		if options.HasDefault && isZeroValue && value.CanAddr() && value.CanInterface() {
			if err := loadYAML([]byte(options.DefaultValue), value.Addr().Interface()); err != nil {
				return fmt.Errorf("invalid default value for \"%s\": %w", field.Name, err)
			}
		}
	}
	return nil
}
