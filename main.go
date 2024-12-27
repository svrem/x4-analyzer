package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Station struct {
	ID    string
	Code  string
	Owner string
	Name  string
}

type StationPageData struct {
	Stations []Station
}

type TradeOffer struct {
	Offertype string
	Ware      string
	Price     float32
	Amount    float32
}

type TradeOfferPageData struct {
	Station     Station
	TradeOffers []TradeOffer
}

func main() {
	println("Hello, World!")
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/stations", func(w http.ResponseWriter, r *http.Request) {
		stationsTemplate := template.Must(template.ParseFiles("views/stations.html"))

		stations, err := db.Query("SELECT * FROM stations")
		if err != nil {
			log.Fatal(err)
		}
		defer stations.Close()

		var stationList []Station
		for stations.Next() {
			var id string
			var code string
			var owner string
			var name string
			stations.Scan(&id, &code, &owner, &name)

			station := Station{
				ID:    id,
				Code:  code,
				Owner: owner,
				Name:  name,
			}
			stationList = append(stationList, station)
		}

		data := StationPageData{
			Stations: stationList,
		}

		stationsTemplate.Execute(w, data)
	})

	http.Handle("/stations/{id}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		stationTemplate := template.Must(template.ParseFiles("views/station.html"))

		id := r.PathValue("id")

		stations, err := db.Query("SELECT code, owner FROM stations WHERE id = ?", id)
		if err != nil {
			log.Fatal(err)
		}
		defer stations.Close()

		if !stations.Next() {
			http.NotFound(w, r)
			return
		}

		var code string
		var owner string
		stations.Scan(&code, &owner)

		station := Station{
			ID:    id,
			Code:  code,
			Owner: owner}

		fmt.Println("Station: ", station)

		tradeOffers, err := db.Query("SELECT * FROM tradeoffers WHERE station_id = ?", id)
		if err != nil {
			log.Fatal(err)
		}
		defer tradeOffers.Close()

		var tradeOfferList []TradeOffer
		for tradeOffers.Next() {
			var offertype string
			var ware string
			var price float32
			var amount float32
			var station_id string

			tradeOffers.Scan(&offertype, &ware, &price, &amount, &station_id)

			tradeOffer := TradeOffer{
				Offertype: offertype,
				Ware:      ware,
				Price:     price,
				Amount:    amount}
			tradeOfferList = append(tradeOfferList, tradeOffer)

		}

		data := TradeOfferPageData{
			Station:     station,
			TradeOffers: tradeOfferList,
		}

		fmt.Println("TradeOffers: ", tradeOfferList)

		stationTemplate.Execute(w, data)
	}))

	slog.Info("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
