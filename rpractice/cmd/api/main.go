package main

import (
	"fmt"
	"os"
	"os/signal"

	"rpractice/internal/controller/http/router"
	"rpractice/internal/usecase"
	"rpractice/pkg/httpServer"
)

func main() {
	useCase := usecase.NewUserUseCase()
	rout := router.NewRouter(useCase)

	httpServ := httpServer.NewServer(rout)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		fmt.Printf("Got signal: %s\n", s.String())

	case err := <-httpServ.Notify():
		fmt.Printf("app - Run - httpServer.Notify(): %s\n", err.Error())
	}

	if err := httpServ.Shutdown(); err != nil {
		fmt.Printf("app - Run - httpServer.Shutdown(): %s\n", err.Error())
	}
}
