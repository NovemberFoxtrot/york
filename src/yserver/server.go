package yserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ytemplate"
	"os"
	"yindex"
	"sir"
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

type Response map[string]interface{}

func (r Response) String() string {
	b, err := json.Marshal(r)

	sir.CheckError(err)

	return string(b)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, Response{"success": true, "message": "Hello!"})
	return
}

func New() {
	ytemplate.ThePool.Fill("index", "templates/layout.html", "templates/index.html")
	ytemplate.ThePool.Fill("search", "templates/layout.html", "templates/search.html")

	wd, err := os.Getwd()
	sir.CheckError(err)

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/add", AddHandler)
	http.HandleFunc("/test", TestHandler)
	http.HandleFunc("/search", SearchHandler)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(wd+`/public`))))

	err = http.ListenAndServe(":9090", nil)
	sir.CheckError(err)
}
