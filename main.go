package main

import (
	"database/sql"
	"log"
	"log/slog"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/svrem/x4-analyzer/internal/views"
)

func main() {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/stations/", func(w http.ResponseWriter, r *http.Request) {
		views.HandleStationsPage(w, r, db)
	})

	http.Handle("/stations/{id}/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		views.HandleIndividualStationPage(w, r, db)
	}))

	http.HandleFunc("/trades/", func(w http.ResponseWriter, r *http.Request) {
		views.HandleBestTradeOptionsPage(w, r, db)
	})

	http.HandleFunc("/", views.HandleIndexPage)

	slog.Info("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
