package main

import (
	"net/http"

	"github.com/o-mercan/Wallet-Service-Api/internal/database"
	"github.com/o-mercan/Wallet-Service-Api/internal/model"
	transportHTTP "github.com/o-mercan/Wallet-Service-Api/internal/transport/http"
	log "github.com/sirupsen/logrus"
)

// App - the struct with contains like pointers
// database connections
type App struct{}

func (a *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Setting up our app")

	db, err := database.NewDatabase()
	if err != nil {
		log.Error("Migrate Error")
	}

	if err = database.MigrateDB(db); err != nil {
		log.Error("Migrate error")
		return err
	}

	transactionService := model.NewTransactionService(db)
	//userService := model.NewUserService(db)
	//walletService := model.NewWalletService(db)

	handler := transportHTTP.NewHandler(transactionService)
	handler.SetUpRoutes()

	if err := http.ListenAndServe(":9000", handler.Router); err != nil {
		log.Error("Failed to set up server")
		return err
	}
	log.Info("Server is running on port 9000")
	return nil

}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		log.Error("Error starting API")
	}

}
