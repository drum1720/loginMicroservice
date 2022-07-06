package main

import (
	"loginMicroservice/app/internal/application"
	"loginMicroservice/app/pkg/configs"
)

func main() {
	configs.LoadEnvConfigs()

	loginServer := application.NewLoginServer()
	loginServer.Run()

	//проверка на  что-то, и перезапуск, остановка
}
