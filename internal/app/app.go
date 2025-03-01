package app

import (
	"fmt"
	"log"
	"university-db-admin/internal/config"
	"university-db-admin/pkg/dbclient"
)

type App struct {
	cfg *config.Config
}

func NewApp() App {
	log.Println("initializing application")

	log.Println("initializing config")
	cfg := config.LoadConfig()
	fmt.Println(cfg.DB)

	log.Println("initializing database")
	pgClient := dbclient.NewClient(cfg.DB)
	log.Println(pgClient)

	log.Println("application initialized")

	return App{
		cfg: cfg,
	}
}

func (a *App) startUI() {
	// TODO
}

func Run() {
	app := NewApp()

	log.Println("starting application")
	app.startUI()
}
