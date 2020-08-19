package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var (
	config *Config
	cred   map[string]string
)

const (
	envDevelopment = "development"
	envStaging     = "jx-staging"
	envProduction  = "jx-production"
)

type option struct {
	configFile      string
	credentialsFile string
}

// Init ...
func Init(opts ...Option) error {
	opt := &option{
		configFile:      getDefaultConfigFile(),
		credentialsFile: getCredentialsFile(),
	}
	for _, optFunc := range opts {
		optFunc(opt)
	}

	out, err := ioutil.ReadFile(opt.configFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(out, &config)
	if err != nil {
		return err
	}

	credentials, err := ioutil.ReadFile(opt.credentialsFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(credentials, &cred)
	if err != nil {
		return err
	}

	return nil
}

// Option ...
type Option func(*option)

// WithConfigFile ...
func WithConfigFile(file string) Option {
	return func(opt *option) {
		opt.configFile = file
	}
}

func getDefaultConfigFile() string {
	var (
		repoPath     = filepath.Join(os.Getenv("GOPATH"), "src/firebase")
		configPath   = filepath.Join(repoPath, "files/etc/firebase/firebase.development.yaml")
		namespace, _ = ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	)

	env := string(namespace)
	if os.Getenv("GOPATH") == "" {
		configPath = "files/etc/firebase/firebase.development.yaml"
	}

	if env != "" {
		if env == envStaging {
			configPath = "./firebase.staging.yaml"
		} else if env == envProduction {
			configPath = "./firebase.production.yaml"
		}
	}
	return configPath
}

// // Get ...
// func Get() *Config {
// 	return config
// }

func getCredentialsFile() string {
	var (
		repoPath     = filepath.Join(os.Getenv("GOPATH"), "src/firebase")
		configPath   = filepath.Join(repoPath, "files/etc/firebase/credentials.development.json")
		namespace, _ = ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	)

	if os.Getenv("GOPATH") == "" {
		configPath = "files/etc/firebase/credentials.development.json"
	}

	env := string(namespace)

	if env != "" {
		if env == envStaging {
			configPath = "./credentials.staging.json"
		}
		if env == envProduction {
			configPath = "/credentials.production.json"
		}
	}
	return configPath
}

// Get ...
func Get() (*Config, map[string]string) {
	return config, cred
}
