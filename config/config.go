package config

import (
	"log"

	"github.com/spf13/viper"
)

type (
	Proxy struct {
		URL         string `mapstructure:"url"`
		CountryCode string `mapstructure:"countryCode"`
	}

	AppCfg struct {
		BotToken string  `mapstructure:"tg_bot_token"`
		Proxies  []Proxy `mapstructure:"proxies"`
	}
)

func MustLoadConfig(cfgPath string) (cfg AppCfg) {
	viper.SetConfigFile(cfgPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("can't read config from path %s", cfgPath)
		return
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("can't marshal config, read from path %s", cfgPath)
		return
	}

	return
}
