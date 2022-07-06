package core

import (
	"errors"
	"loginMicroservice/app/pkg/configs"
)

type Cfg struct {
	dsnPg string
	url   string
}

func InitCfg() (*Cfg, error) {
	dsnPg := configs.GetEnvDefault("DSN", "")
	address := configs.GetEnvDefault("URL", "")
	if dsnPg == "" || address == "" {
		return nil, errors.New("env config is empty")
	}

	return &Cfg{dsnPg: dsnPg, url: address}, nil
}

func (c Cfg) GetDsnPG() string {
	return c.dsnPg
}

func (c Cfg) GetUrl() string {
	return c.url
}
