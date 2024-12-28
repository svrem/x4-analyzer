package views

import (
	"net/http"
	"text/template"
)

// func getTmpl() map[string]*template.Template {
// 	tmpl := make(map[string]*template.Template)

// 	tmpl["index"] = template.Must(template.ParseFiles("views/index.html", "views/components/header.html"))
// 	tmpl["stations"] = template.Must(template.ParseFiles("views/stations.html", "views/components/header.html"))
// 	tmpl["station"] = template.Must(template.ParseFiles("views/station.html", "views/components/header.html"))
// 	tmpl["trades"] = template.Must(template.ParseFiles("views/trades.html"))

// 	return tmpl
// }

type PageData struct {
	Title string
}

func HandleIndexPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	renderTemplate(w, r, "views/index.html", PageData{
		Title: "Index",
	})
}

func renderTemplate(w http.ResponseWriter, r *http.Request, tmpl_name string, data interface{}) {
	condensed := r.URL.Query().Get("c")

	var tmpl *template.Template

	if condensed == "true" {
		tmpl = template.Must(template.ParseFiles("views/util/content.html", tmpl_name))
	} else {
		tmpl = template.Must(template.ParseFiles("views/components/base.html", tmpl_name, "views/components/header.html", "views/components/sidebar.html"))
	}

	tmpl.Execute(w, data)
}
