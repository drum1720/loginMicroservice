package main

import (
	"loginMicroservice/app/internal/loginService"
	"loginMicroservice/app/pkg/configs"
)

func main() {
	configs.LoadEnvConfigs()

	loginServer := loginService.NewLoginServer()
	loginServer.Init()
	loginServer.Run()
}
