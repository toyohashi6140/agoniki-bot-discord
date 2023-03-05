package config

import "github.com/kelseyhightower/envconfig"

var DiscordEnv discordEnv

type discordEnv struct {
	ClientID   string   `envconfig:"CLIENT_ID" require:"true"`
	Token      string   `envconfig:"BOT_TOKEN" require:"true"`
	Swearings  []string `envconfig:"SWEARINGS" require:"true"`
	Swearings2 []string `envconfig:"SWEARINGS2" require:"true"`
}

func init() {
	if err := envconfig.Process("", &DiscordEnv); err != nil {
		panic(err.Error())
	}
}
