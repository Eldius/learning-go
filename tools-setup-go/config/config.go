/*
Package config groups the config related code
*/
package config

import (
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
	"os"
)

const (
	configFolder = "~/.tools-setup"
)

func init() {
	if cfgDir, err := homedir.Expand(configFolder); err != nil {
		panic(err.Error())
	} else {
		os.MkdirAll(cfgDir, os.ModePerm)
	}
}

// AppConfig is an abstractio for app config parameters
type AppConfig struct {
	LogFolder string
	DbFolder  string
	BinFolder string
	Verbose   bool
}

// LogFolderExpanded Returns log folder
func (a *AppConfig) LogFolderExpanded() string {
	if folder, err := homedir.Expand(a.LogFolder); err != nil {
		panic(err.Error())
	} else {
		return folder
	}
}

// BinFolderExpanded Returns bin folder
func (a *AppConfig) BinFolderExpanded() string {
	if folder, err := homedir.Expand(a.BinFolder); err != nil {
		panic(err.Error())
	} else {
		return folder
	}
}

// DbFolderExpanded returns db folder
func (a *AppConfig) DbFolderExpanded() string {
	if folder, err := homedir.Expand(a.DbFolder); err != nil {
		panic(err.Error())
	} else {
		return folder
	}
}
// LoadConfig loads configuration file
func LoadConfig() AppConfig {
	var cfg AppConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err.Error())
	}

	debugCfg(cfg)
	return cfg
}

func debugCfg(cfg AppConfig) {
	if cfg.Verbose {
		bs, err := yaml.Marshal(cfg)
		if err != nil {
			//log.Fatalf("unable to marshal config to YAML: %v", err)
			panic(err.Error())
		}
		fmt.Println(fmt.Sprintf("---\n%s\n---", string(bs)))
	}
}