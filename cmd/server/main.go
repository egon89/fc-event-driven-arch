package main

import (
	"context"
	"log"
	"net/http"

	"github.com/egon89/fc-event-driven-arch/internal/config"
	"github.com/egon89/fc-event-driven-arch/internal/db"
	"github.com/egon89/fc-event-driven-arch/internal/healthz"
	"github.com/egon89/fc-event-driven-arch/internal/kafka"
	"github.com/egon89/fc-event-driven-arch/internal/repository"
	"github.com/egon89/fc-event-driven-arch/internal/service"
	"github.com/egon89/fc-event-driven-arch/internal/usecase"
	"github.com/go-chi/chi"
)

func main() {
	cfg := config.Load()

	db, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	ctx := context.Background()

	balanceRepository := repository.NewRepository(db)

	balanceService := service.NewBalanceService(balanceRepository)

	saveBalanceUseCase := usecase.NewSaveBalanceUseCase(balanceService)

	balanceConsumer := kafka.NewBalanceConsumer(cfg.KafkaBroker, cfg.KafkaTopic, saveBalanceUseCase)

	go balanceConsumer.Consume(ctx)

	healthzHandler := healthz.NewHealthzHandler()

	r := chi.NewRouter()
	r.Mount("/healthz", healthzHandler.Routes())

	log.Printf("Server running on port %s", cfg.AppPort)

	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, r))
}
