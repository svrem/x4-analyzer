package views

import (
	"net/http"
	"text/template"
)

func getTmpl() map[string]*template.Template {
	tmpl := make(map[string]*template.Template)

	tmpl["index"] = template.Must(template.ParseFiles("views/index.html", "views/components/header.html"))
	tmpl["stations"] = template.Must(template.ParseFiles("views/stations.html", "views/components/header.html"))
	tmpl["station"] = template.Must(template.ParseFiles("views/station.html", "views/components/header.html"))
	tmpl["trades"] = template.Must(template.ParseFiles("views/trades.html"))

	return tmpl
}

func HandleIndexPage(w http.ResponseWriter, r *http.Request) {
	tmpl := getTmpl()

	tmpl["index"].Execute(w, nil)
}
