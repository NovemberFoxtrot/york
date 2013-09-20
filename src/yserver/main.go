package yserver

import (
	"net/http"
	"os"
	"yeasy"
	"yindex"
	"ytemplate"
	"ytext"
)

type tv struct {
	Location string
	GramsLen int
}

type stv struct {
	Result yindex.QueryResult
	Score  int
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	go yindex.Add(r.FormValue("website"))
	http.Redirect(w, r, "/", http.StatusFound)
}

func MediaHandler(w http.ResponseWriter, r *http.Request) {
	ytemplate.ThePool.Fill("media", "templates/layout.html", "templates/media.html")
	ytemplate.ThePool.Pools["media"].Execute(w, nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/magic/"+r.URL.String()[len(`/images/`):])
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	ytemplate.ThePool.Fill("search", "templates/layout.html", "templates/search.html")

	stvs := make([]stv, 0)

	for index, location := range yindex.Query(r.FormValue("query")) {
		stvs = append(stvs, stv{location, index})
	}

	ytemplate.ThePool.Pools["search"].Execute(w, stvs)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	ytemplate.ThePool.Fill("index", "templates/layout.html", "templates/index.html")

	tvs := make([]tv, 0)
	tvs = append(tvs, tv{"Index", yindex.IndexDataLen()})

	for i := 0; i < len(ytext.TheDocuments); i++ {
		tvs = append(tvs, tv{ytext.TheDocuments[i].Location, len(ytext.TheDocuments[i].Grams)})
	}

	ytemplate.ThePool.Pools["index"].Execute(w, tvs)
}

func New() {
	wd, err := os.Getwd()
	yeasy.CheckError(err)

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/add", AddHandler)
	http.HandleFunc("/images/", HomeHandler)
	http.HandleFunc("/media", MediaHandler)
	http.HandleFunc("/search", SearchHandler)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(wd+`/public`))))

	err = http.ListenAndServe(":6969", nil)
	yeasy.CheckError(err)
}
