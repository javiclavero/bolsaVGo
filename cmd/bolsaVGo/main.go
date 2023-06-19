// main.go
package main

import (
	"bolsaVGo/config"
	"bolsaVGo/pkg/controllers"
	"bolsaVGo/pkg/middleware"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/bolsaVGo", config.DBUser, config.DBPassword))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tc := &controllers.TickerController{DB: db}

	r := mux.NewRouter()

	// Ruta para index.html
	r.HandleFunc("/bolsaVGo", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html")
	})

	// Ruta protegida con autenticación específica para tickers
	r.HandleFunc("/tickers", middleware.AuthMiddleware(tc.GetTickers)).Methods("GET")
	//r.Handle("/tickers", middleware.AuthMiddleware(tc.GetTickers)).Methods("GET")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
