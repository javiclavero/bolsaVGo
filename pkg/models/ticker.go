// models/ticker.go
package models

import (
	"database/sql"
)

type Ticker struct {
	Empresa string
	Ticker  string
}

func GetTickers(db *sql.DB) ([]Ticker, error) {
	rows, err := db.Query("SELECT tic_empresa, tic_ticker FROM tblTickers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tickers := []Ticker{}
	for rows.Next() {
		var t Ticker
		if err := rows.Scan(&t.Empresa, &t.Ticker); err != nil {
			return nil, err
		}
		tickers = append(tickers, t)
	}

	return tickers, nil
}
