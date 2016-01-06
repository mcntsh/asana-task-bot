package config

import (
	"fmt"
	"os"
)

type Config struct {
	Address string
	APIKey  string
}

var Configuration *Config

func ParseConfig() {
	Configuration = &Config{}

	Configuration.Address = fmt.Sprintf(":%v", os.Getenv("PORT"))
	Configuration.APIKey = "xoxp-2182767778-3338739518-17735596101-6d474c4071"

	return
}
