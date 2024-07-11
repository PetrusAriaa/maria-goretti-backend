package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/PetrusAriaa/web-margot-backend/lib/controller"
	"github.com/PetrusAriaa/web-margot-backend/lib/db"
	"github.com/PetrusAriaa/web-margot-backend/lib/service"
)

var s *http.Server

func main() {
	err := godotenv.Load(".env.development")
	if err != nil {
		log.Default().Fatal(err.Error())
	}

	ctx := context.Background()

	client := db.NewDBConnections(ctx)
	defer client.CloudStorageClient.Close()

	srv := service.NewService(client, ctx)
	c := controller.NewController(srv)

	mux := http.NewServeMux()

	s = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	mux.HandleFunc("/api/v1/images", c.GetImageList)
	mux.HandleFunc("/assets/images/{id}", c.GetImage)

	log.Default().Printf("Server started on http://0.0.0.0:8080")
	if err := s.ListenAndServe(); err != nil {
		client.CloudStorageClient.Close()
		panic(err.Error())
	}
}

func StartServer(w http.ResponseWriter, r *http.Request) {
	main()
	s.ListenAndServe()
}
