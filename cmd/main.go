package main

import (
	"fmt"
	"log"

	"github.com/adrianosiqe/eulabs-challenge-api/internal/domains/repositories"
	server "github.com/adrianosiqe/eulabs-challenge-api/internal/http"
	"github.com/adrianosiqe/eulabs-challenge-api/pkg/config"
	"github.com/adrianosiqe/eulabs-challenge-api/pkg/database"
)

func main() {
	config.LoadConfig()

	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	productRepository := repositories.NewProductRepository(db)

	address := fmt.Sprintf(":%s", config.Cfg.PORT)
	http := server.NewServer(productRepository)
	http.RouteInit(address)
}
