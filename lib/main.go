package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	h1 "github.com/PetrusAriaa/web-margot-backend/api"
	h2 "github.com/PetrusAriaa/web-margot-backend/api/static"
)

func main() {
	godotenv.Load(".env.development")

	mux := http.NewServeMux()

	s := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	mux.HandleFunc("/api/images", h1.Image)
	mux.HandleFunc("/api/status", h1.Status)
	mux.HandleFunc("/api/static/img", h2.ImageHandler)

	log.Default().Printf("Server started on http://0.0.0.0:8080")
	if err := s.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}
