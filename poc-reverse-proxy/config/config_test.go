package config

import (
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/spf13/viper"
)

func init() {
	if _, err := os.Stat("./samples/config.yaml"); err != nil {
		log.Panicf("Failed to stat file: %s\n", err.Error())
	}
	viper.SetConfigFile("samples/config.yaml")
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func TestLoadRoutes(t *testing.T) {
	if cfg, err := LoadRoutes(); err != nil {
		t.Error("Failed to load configuration from file")
	} else {
		log.Println(cfg)
		if len(cfg.Routes) != 2 {
			log.Println(cfg)
			t.Errorf("Should have 2 routes, but has %d", len(cfg.Routes))
		}
	}
}

func TestMatch(t *testing.T) {
	if cfg, err := LoadRoutes(); err != nil {
		t.Error("Failed to load configuration from file")
	} else {

		r := match("/app01", cfg.patterns)
		if r == nil {
			t.Error("Should return a non nil route")
		} else if r.Backends[0] != "http://localhost:8888" {
			t.Errorf("Backend should be 'http://localhost:8888', but was '%s'", r.Backends[0])
		}
	}
}

func TestNotMatch(t *testing.T) {
	if cfg, err := LoadRoutes(); err != nil {
		t.Error("Failed to load configuration from file")
	} else {

		r := match("/app", cfg.patterns)
		if r != nil {
			t.Error("Should return a nil route")
		}
	}
}

func TestMatchPrefix(t *testing.T) {
	if cfg, err := LoadRoutes(); err != nil {
		t.Error("Failed to load configuration from file")
	} else {
		r := match("/app01/xpto", cfg.patterns)
		if r.Backends[0] != "http://localhost:8888" {
			t.Errorf("Backend should be 'http://localhost:8888', but was '%s'", r.Backends[0])
		}
	}
}

func TestDraft(t *testing.T) {
	pattern01 := `^/app01.*\z`
	r := regexp.MustCompile(pattern01)
	if !r.MatchString("/app01") {
		t.Errorf("Failed to match pattern01 with /app01")
	}
	if !r.MatchString("/app01/xpto") {
		t.Errorf("Failed to match pattern01 with /app01/xpto")
	}
}

func TestDraft2(t *testing.T) {
	for k, v := range viper.GetStringMap("routes") {
		log.Println(k, "=>", v)
	}
}
