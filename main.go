package main

import (
	"errors"
	"golnfuturecapacities/api/config/database"
	"golnfuturecapacities/api/models"
	"golnfuturecapacities/api/models/products"
	"golnfuturecapacities/api/server"
	"log"
	"net/http"
	"os"
)

var (
	CD = database.CDriver()
)

func main() {
	models.AutoMigration(CD)
	products.ProductMigration(CD)
	s := server.NewAPIServer(":8080")
	if err := s.Run(); err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
		os.Exit(1)
	}
}
