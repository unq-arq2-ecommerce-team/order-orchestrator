package config

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		Environment string            `required:"true" default:"development"`
		Port        int               `required:"true" default:"8080"`
		LogLevel    string            `split_words:"true" default:"DEBUG"`
		OrderUrl    OrderEndpoints    `split_words:"true" required:"true"`
		CustomerUrl CustomerEndpoints `split_words:"true" required:"true"`
	}
	OrderEndpoints struct {
		FindById  string `split_words:"true" required:"true"`
		Create    string `split_words:"true" required:"true"`
		Confirm   string `split_words:"true" required:"true"`
		Delivered string `split_words:"true" required:"true"`
	}
	CustomerEndpoints struct {
		FindById string `split_words:"true" required:"true"`
	}
)

func LoadConfig() Config {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		panic(err.Error())
	}
	return config
}
