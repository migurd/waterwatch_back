package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/migurd/waterwatch_back/controllers"
	"github.com/migurd/waterwatch_back/database"
	"github.com/migurd/waterwatch_back/models"
	"github.com/migurd/waterwatch_back/router"
)

type Config struct {
	Port string
	Env  string // Environment variable to differentiate between production and development
}

type Application struct {
	Config      Config
	Models      models.Models
	Controllers controllers.Controllers
}

var (
	port = os.Getenv("PORT")
	env  = os.Getenv("ENV") // Environment variable to check if it's production
)

func (app *Application) Serve() error {
	fmt.Println("API listening on port", app.Config.Port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.Config.Port),
		Handler: router.Routes(&app.Controllers),
	}

	if app.Config.Env == "production" {
		// Use HTTPS in production with the Let's Encrypt certificate
		return srv.ListenAndServeTLS("/etc/letsencrypt/live/angelqui.xyz/fullchain.pem", "/etc/letsencrypt/live/angelqui.xyz/privkey.pem")
	} else {
		// Use HTTP for development
		return srv.ListenAndServe()
	}
}

func main() {
	var cfg Config
	cfg.Port = port
	cfg.Env = env

	dsn := os.Getenv("DSN")
	dbConn, err := database.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Database couldn't be opened", err)
	}

	defer dbConn.DB.Close()

	models := models.New(dbConn.DB)
	controllers := controllers.New(dbConn.DB)

	app := &Application{
		Config:      cfg,
		Models:      models,
		Controllers: controllers,
	}

	err = app.Serve()
	if err != nil {
		log.Fatal("Error starting server", err)
	}
}
