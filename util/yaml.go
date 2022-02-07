package util

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type ElasticsearchInfo struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type GCSInfo struct {
	Bucket string `yaml:"bucket"`
}

type TokenInfo struct {
	Secret string `yaml:"secret"`
}

type ApplicationConfig struct {
	ElasticsearchConfig *ElasticsearchInfo `yaml:"elasticsearch"`
	GCSConfig           *GCSInfo           `yaml:"gcs"`
	TokenConfig         *TokenInfo         `yaml:"token"`
}

func LoadApplicationConfig(configDir, configFile string) (*ApplicationConfig, error) {
	content, err := ioutil.ReadFile(filepath.Join(configDir, configFile))
	if err != nil {
		return nil, err
	}

	var config ApplicationConfig
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
