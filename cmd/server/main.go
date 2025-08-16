package main

import (
	"log"
	"net/http"

	"github.com/egon89/fc-event-driven-arch/internal/config"
	"github.com/egon89/fc-event-driven-arch/internal/db"
	"github.com/egon89/fc-event-driven-arch/internal/healthz"
	"github.com/egon89/fc-event-driven-arch/internal/kafka"
	"github.com/go-chi/chi"
)

func main() {
	cfg := config.Load()

	_, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	go kafka.Consumer(cfg.KafkaBroker, cfg.KafkaTopic)

	healthzHandler := healthz.NewHealthzHandler()

	r := chi.NewRouter()
	r.Mount("/healthz", healthzHandler.Routes())

	log.Printf("Server running on port %s", cfg.AppPort)

	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, r))
}
