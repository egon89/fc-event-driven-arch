package web

import (
	"encoding/json"
	"net/http"

	"github.com/egon89/fc-event-driven-arch/internal/usecase"
	"github.com/go-chi/chi"
)

type FindBalanceByAccountIdResponseDto struct {
	Id        string  `json:"id"`
	AccountId string  `json:"accountId"`
	Balance   float64 `json:"balance"`
}

type WebBalanceHandler struct {
	FindBalanceByAccountIdUseCase usecase.FindBalanceByAccountIdUseCase
}

func NewWebBalanceHandler(
	findBalanceByAccountIdUseCase usecase.FindBalanceByAccountIdUseCase,
) *WebBalanceHandler {
	return &WebBalanceHandler{
		FindBalanceByAccountIdUseCase: findBalanceByAccountIdUseCase,
	}
}

func (h *WebBalanceHandler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/{accountId}", h.FindBalanceByAccountId)

	return r
}

func (h *WebBalanceHandler) FindBalanceByAccountId(w http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "accountId")

	balance, err := h.FindBalanceByAccountIdUseCase.Execute(r.Context(), accountId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := FindBalanceByAccountIdResponseDto{
		Id:        balance.Id,
		AccountId: balance.AccountId,
		Balance:   balance.Balance,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
