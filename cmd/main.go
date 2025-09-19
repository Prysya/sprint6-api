package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags|log.Lshortfile)

	srv := server.CreateServer(logger)

	if err := srv.Start(); err != nil {
		logger.Fatalf("Ошибка запуска сервера: %s", err.Error())
	}
}
