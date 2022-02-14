package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

type testConfig struct {
	Host string
	Port int
	Mode string
	SSL  bool
}

func newTestConfig() testConfig {
	return testConfig{
		Host: "127.0.0.1",
		Port: 80,
		Mode: "debug",
		SSL:  true,
	}
}

func TestLoadJYAML(t *testing.T) {
	file, err := ioutil.TempFile("", "config-*.yaml")
	if err != nil {
		t.Errorf("failed to create yaml config file")
	}
	defer file.Close()
	defer os.Remove(file.Name())

	cfg := newTestConfig()
	bytes, err := yaml.Marshal(cfg)
	if err != nil {
		t.Errorf("failed to marshal yaml config")
	}
	file.Write(bytes)

	var result testConfig
	err = Load(file.Name(), &result)
	if err != nil {
		t.Errorf("config load failed: %v", err)
	}

	if !reflect.DeepEqual(result, cfg) {
		t.Errorf("bad load result: %v != %v", result, cfg)
	}
}

func TestLoadJTOML(t *testing.T) {
	file, err := ioutil.TempFile("", "config-*.toml")
	if err != nil {
		t.Errorf("failed to create toml config file")
	}
	defer file.Close()
	defer os.Remove(file.Name())

	cfg := newTestConfig()
	var buffer bytes.Buffer
	if err := toml.NewEncoder(&buffer).Encode(cfg); err != nil {
		t.Errorf("failed to encode toml config")
	}
	file.Write(buffer.Bytes())

	var result testConfig
	err = Load(file.Name(), &result)
	if err != nil {
		t.Errorf("config load failed: %v", err)
	}

	if !reflect.DeepEqual(result, cfg) {
		t.Errorf("bad load result: %v != %v", result, cfg)
	}
}

func TestLoadJSON(t *testing.T) {
	file, err := ioutil.TempFile("", "config-*.json")
	if err != nil {
		t.Errorf("failed to create json config file")
	}
	defer file.Close()
	defer os.Remove(file.Name())

	cfg := newTestConfig()
	bytes, err := json.Marshal(cfg)
	if err != nil {
		t.Errorf("failed to marshal json config")
	}
	file.Write(bytes)

	var result testConfig
	err = Load(file.Name(), &result)
	if err != nil {
		t.Errorf("config load failed: %v", err)
	}

	if !reflect.DeepEqual(result, cfg) {
		t.Errorf("bad load result: %v != %v", result, cfg)
	}
}

type testDefaultConfig struct {
	Host string
	Port int `config:"default=80"`
	Mode string
	SSL  bool `config:"default=true"`
}

func TestLoadDefault(t *testing.T) {
	file, err := ioutil.TempFile("", "config-*.json")
	if err != nil {
		t.Errorf("failed to create json config file")
	}
	defer file.Close()
	defer os.Remove(file.Name())

	cfg := testDefaultConfig{
		Host: "127.0.0.1",
		Mode: "debug",
	}
	bytes, err := json.Marshal(cfg)
	if err != nil {
		t.Errorf("failed to marshal json config")
	}
	file.Write(bytes)

	var result testDefaultConfig
	err = Load(file.Name(), &result)
	if err != nil {
		t.Errorf("config load failed: %v", err)
	}

	if result.Host != cfg.Host || result.Mode != cfg.Mode {
		t.Errorf("bad load result: %v", result)
	}
	if result.Port != 80 || result.SSL != true {
		t.Errorf("bad default value: %v", result)
	}
}

type testRequiredConfig struct {
	Host string `config:"required"`
	Port int    `config:"required"`
	Mode string
	SSL  bool
}

func TestLoadRequired(t *testing.T) {
	file, err := ioutil.TempFile("", "config-*.json")
	if err != nil {
		t.Errorf("failed to create json config file")
	}
	defer file.Close()
	defer os.Remove(file.Name())

	cfg := testRequiredConfig{
		Port: 8000,
		Mode: "debug",
	}
	bytes, err := json.Marshal(cfg)
	if err != nil {
		t.Errorf("failed to marshal json config")
	}
	file.Write(bytes)

	var result testRequiredConfig
	err = Load(file.Name(), &result)
	if err == nil {
		t.Errorf("required validation failed")
	}
}
