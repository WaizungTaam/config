package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
)

var (
	defaultTagName   = "config"
	defaultDelimeter = ";"
)

// Loader is a config loader
type Loader struct {
	options *Options
}

// Options contains loader options
type Options struct {
	TagName   string
	Delimeter string
}

// New creates a new loader with default options
func New() *Loader {
	return &Loader{
		options: &Options{
			TagName:   defaultTagName,
			Delimeter: defaultDelimeter,
		},
	}
}

// NewWithOptions creates a new loader with given options
func NewWithOptions(options *Options) *Loader {
	return &Loader{options: options}
}

// With applies options
func (d *Loader) With(f func(*Options)) *Loader {
	f(d.options)
	return d
}

// Load loads config from a file
func (d Loader) Load(file string, config interface{}) error {
	if err := d.read(file, config); err != nil {
		return err
	}
	if err := d.process(config); err != nil {
		return err
	}
	return nil
}

func (d Loader) read(file string, config interface{}) error {
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

func (d Loader) process(config interface{}) error {
	configValue := reflect.Indirect(reflect.ValueOf(config))
	if configValue.Kind() != reflect.Struct {
		return fmt.Errorf("struct required")
	}
	configType := configValue.Type()
	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		value := configValue.Field(i)

		tag := field.Tag.Get(d.options.TagName)
		options, err := parseTag(tag, d.options.Delimeter)
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

// Load loads config from a file with default options
func Load(file string, config interface{}) error {
	return New().Load(file, config)
}
