package main

import (
	"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/PetrusAriaa/web-margot-backend/src/controller"
	"github.com/PetrusAriaa/web-margot-backend/src/db"
	"github.com/PetrusAriaa/web-margot-backend/src/service"
)

func main() {
	err := godotenv.Load(".env.development")
	if err != nil {
		log.Default().Fatal(err.Error())
	}

	ctx := context.Background()

	client := db.NewDBConnections(ctx)
	srv := service.NewService(client, ctx)
	c := controller.NewController(srv)

	mux := http.NewServeMux()

	s := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	mux.HandleFunc("/assets/images", c.GetImageList)
	mux.HandleFunc("/assets/images/{id}", c.GetImage)

	log.Default().Printf("Server started on http://0.0.0.0:8080")
	if err := s.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}
