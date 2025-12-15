package main

import (
	"log"

	"migrant-id/configs"
	"migrant-id/internal/http/handler"
	"migrant-id/internal/http/router"
	"migrant-id/internal/repository"
	"migrant-id/internal/service"
)

func main() {
	log.Println("Starting...")

	cfg := configs.Load()

	repo := repository.NewInMemoryRepo()
	svc := service.NewMigrantService(repo, cfg.JWTSecret)
	h := handler.NewHandler(svc)
	r := router.NewRouter(h)

	srv := &Server{}
	log.Printf("Server is running on port :%s\n", cfg.Port)
	if err := srv.Run(cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}
