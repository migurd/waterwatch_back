package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/migurd/waterwatch_back/database"
	"github.com/migurd/waterwatch_back/router"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
}

var port = os.Getenv("PORT")

func (app *Application) Serve() error {
	fmt.Println("API listening on port ", app.Config.Port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.Config.Port),
		Handler: router.Routes(),
	}

	return srv.ListenAndServe()
}

func main() {
	var cfg Config
	cfg.Port = port

	dsn := os.Getenv("DSN")
	dbConn, err := database.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Database couldn't be opened", err)
	}

	defer dbConn.DB.Close()

	app := &Application{
		Config: cfg,
	}

	err = app.Serve()
	if err != nil {
		log.Fatal("Error creating app", err)
	}
}
