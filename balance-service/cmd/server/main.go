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
	"github.com/egon89/fc-event-driven-arch/internal/web"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
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

	findBalanceByAccountIdUseCase := usecase.NewFindBalanceByAccountIdUseCase(balanceService)
	saveBalanceUseCase := usecase.NewSaveBalanceUseCase(balanceService)

	balanceConsumer := kafka.NewBalanceConsumer(cfg.KafkaBroker, cfg.KafkaTopic, saveBalanceUseCase)

	go balanceConsumer.Consume(ctx)

	healthzHandler := healthz.NewHealthzHandler()
	webBalanceHandler := web.NewWebBalanceHandler(*findBalanceByAccountIdUseCase)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/healthz", healthzHandler.Routes())
	r.Mount("/balances", webBalanceHandler.Routes())

	log.Printf("Server running on port %s", cfg.AppPort)

	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, r))
}
