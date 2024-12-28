package views

import (
	"database/sql"
	"log"
	"net/http"
)

type Station struct {
	ID    string
	Code  string
	Owner string
	Name  string
}

type StationPageData struct {
	Stations []Station
	Title    string
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
	Title       string
}

func HandleStationsPage(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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
		Title:    "Stations",
	}

	renderTemplate(w, r, "views/stations.html", data)
}

func HandleIndividualStationPage(w http.ResponseWriter, r *http.Request, db *sql.DB) {

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
		Owner: owner,
	}

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

	renderTemplate(w, r, "views/station.html", data)
}
