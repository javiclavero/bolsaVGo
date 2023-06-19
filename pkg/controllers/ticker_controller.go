// controllers/ticker_controller.go
package controllers

import (
	"bolsaVGo/pkg/models"
	"database/sql"
	"encoding/json"
	"net/http"
)

type TickerController struct {
	DB *sql.DB
}

func (tc *TickerController) GetTickers(w http.ResponseWriter, r *http.Request) {
	tickers, err := models.GetTickers(tc.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickers)
}
