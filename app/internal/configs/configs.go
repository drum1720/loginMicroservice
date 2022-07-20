package configs

import "loginMicroservice/app/pkg/configs"

type Cfg struct {
	dsnPg string
	url   string
	jwt   string
}

//InitCfg ...
func InitCfg() (Configure, error) {
	dsnPg, err := configs.GetEnvCfg("DSN")
	if err != nil {
		return nil, err
	}

	address, err := configs.GetEnvCfg("URL")
	if err != nil {
		return nil, err
	}

	jwt, err := configs.GetEnvCfg("SUPERSECRETJWT")
	if err != nil {
		return nil, err
	}

	return &Cfg{dsnPg: dsnPg, url: address, jwt: jwt}, nil
}

//GetDsnPG ...
func (c Cfg) GetDsnPG() string {
	return c.dsnPg
}

//GetUrl ...
func (c Cfg) GetUrl() string {
	return c.url
}

//GetJWT ...
func (c Cfg) GetKeyJWT() string {
	return c.jwt
}
