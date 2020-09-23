package config

import (
	"math/rand"
	"log"
	"os"
	"testing"

	"github.com/spf13/viper"
)


func init() {
	if _, err := os.Stat("./samples/config.yaml"); err != nil {
		log.Panicf("Failed to stat file: %s\n", err.Error())
	}
	viper.SetConfigFile("samples/config_benchmark.yaml")
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}

var (
	validPrefixes = []string{"/app01", "/app02", "/app03/xpto", "/app04/abc"}
	invalidPrefixes = []string{"/app", "/xpto", "/app0/xpto", "/abc"}
)

func generateRandomPath(prefix string) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

    b := make([]rune, 10)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
	}
	return prefix + "/" + string(b)
}

func randValidPrefix() string {
	return validPrefixes[rand.Intn(len(validPrefixes))]
}

func randInvalidPrefix() string {
	return invalidPrefixes[rand.Intn(len(invalidPrefixes))]
}

func BenchmarkTest(b *testing.B)  {
	cfg, err := LoadRoutes()
	if err != nil {
		b.Error("Failed to load config")
	}
	
	qtd := b.N
	b.Logf("iterations: %d", qtd)
	for n := 0; n < qtd; n++ {
		if n % 2 == 0 {
			r := match(generateRandomPath(randValidPrefix()), cfg.patterns)
			r.
		} else {
			r := match(generateRandomPath(randInvalidPrefix()), cfg.patterns)
		}
		
	}
}