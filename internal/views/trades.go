package views

import (
	"database/sql"
	"log"
	"net/http"
	"sort"
)

type SellBuyPrice struct {
	Ware string

	MinSellPrice float32
	MaxBuyPrice  float32

	SellerStationID string
	BuyerStationID  string

	TradableVolume float32
}

type TradesPageData struct {
	SellBuyPrices []SellBuyPrice
	Title         string
}

func HandleBestTradeOptionsPage(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	wares, err := db.Query("SELECT DISTINCT ware FROM tradeoffers")
	if err != nil {
		log.Fatal(err)
	}
	defer wares.Close()

	ware_list := []string{}

	for wares.Next() {
		var ware string

		wares.Scan(&ware)
		ware_list = append(ware_list, ware)
	}

	var sellBuyPrices []SellBuyPrice

	for _, ware := range ware_list {
		var maxBuyPrice float32
		var buyerStationID string
		var buyerAmount float32
		if err := db.QueryRow("SELECT MAX(price), station_id, amount FROM tradeoffers WHERE ware = ? AND type = 'buyoffer'", ware).Scan(&maxBuyPrice, &buyerStationID, &buyerAmount); err != nil {
			continue
		}

		var minSellPrice float32
		var sellerStationID string
		var sellerAmount float32
		if err := db.QueryRow("SELECT MIN(price), station_id, amount FROM tradeoffers WHERE ware = ? AND type= 'selloffer'", ware).Scan(&minSellPrice, &sellerStationID, &sellerAmount); err != nil {
			continue
		}

		sellBuyPrices = append(sellBuyPrices, SellBuyPrice{
			Ware: ware,

			MinSellPrice: minSellPrice,
			MaxBuyPrice:  maxBuyPrice,

			SellerStationID: sellerStationID,
			BuyerStationID:  buyerStationID,

			TradableVolume: min(sellerAmount, buyerAmount),
		})
	}

	// sort by profit
	sort.Slice(sellBuyPrices, func(i, j int) bool {
		profitI := sellBuyPrices[i].MaxBuyPrice - sellBuyPrices[i].MinSellPrice
		profitPercantageI := profitI / sellBuyPrices[i].MinSellPrice
		profitJ := sellBuyPrices[j].MaxBuyPrice - sellBuyPrices[j].MinSellPrice
		profitPercantageJ := profitJ / sellBuyPrices[j].MinSellPrice

		maxTradableVolume := max(sellBuyPrices[i].TradableVolume, sellBuyPrices[j].TradableVolume)

		normailzedVolumeI := sellBuyPrices[i].TradableVolume / maxTradableVolume
		normailzedVolumeJ := sellBuyPrices[j].TradableVolume / maxTradableVolume

		return profitPercantageI+normailzedVolumeI > profitPercantageJ+normailzedVolumeJ
	})

	// condensed := r.URL.Query().Get("c")

	// println(condensed)
	// var tradesTemplate *template.Template
	// if condensed == "true" {
	// 	tradesTemplate = template.Must(template.ParseFiles(
	// 		"views/util/content.html",
	// 		"views/trades.html",
	// 	))
	// } else {
	// 	tradesTemplate = template.Must(template.ParseFiles(
	// 		"views/components/base.html",
	// 		"views/trades.html",
	// 		"views/components/header.html",
	// 		"views/components/sidebar.html",
	// 	))
	// }

	// tradesTemplate.Execute(w, TradesPageData{
	// 	SellBuyPrices: sellBuyPrices,
	// })

	renderTemplate(w, r, "views/trades.html", TradesPageData{
		SellBuyPrices: sellBuyPrices,
		Title:         "Best Trade Options",
	})

}
