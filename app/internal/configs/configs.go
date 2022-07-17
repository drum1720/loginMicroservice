package configs

import (
	"loginMicroservice/app/pkg/configs"
)

type Cfg struct {
	dsnPg string
	url   string
}

func (c Cfg) GetDsnPG() string {
	return c.dsnPg
}

func (c Cfg) GetUrl() string {
	return c.url
}

func InitCfg() (*Cfg, error) {
	dsnPg, err := configs.GetEnvCfg("DSN")
	if err != nil {
		return nil, err
	}

	address, err := configs.GetEnvCfg("URL")
	if err != nil {
		return nil, err
	}

	return &Cfg{dsnPg: dsnPg, url: address}, nil
}
