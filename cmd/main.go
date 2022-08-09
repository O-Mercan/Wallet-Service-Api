package main

import (
	log "github.com/sirupsen/logrus"
)
// App - the struct with contains like pointers
// database connections
type App struct{}

func (a *App) Run() error {
	log.SetFormatter(&log.JSONFormatter())
	log.Info("Setting up our app")

	db,err := database.NewDatabase()
	if err != nil{
		log.Error("Migrate Error")
	}

	if err = database



}

func main() {

}
