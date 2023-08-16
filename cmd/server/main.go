package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/divine-within/morsetranslate-go/api/router"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
}

func (app *Application) Serve() error {
	fmt.Println("API is running on port", app.Config.Port)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.Config.Port),
		Handler: router.Routes(),
	}

	return server.ListenAndServe()
}

func main() {
	config := Config{
		Port: "4444",
	}

	app := &Application{
		Config: config,
	}

	err := app.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
