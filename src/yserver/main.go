package yserver

import (
	"net/http"
	"os"
	"yeasy"
	"ytemplate"
)

func MediaHandler(w http.ResponseWriter, r *http.Request) {
	ytemplate.ThePool.Fill("media", "templates/layout.html", "templates/media.html")
	ytemplate.ThePool.Pools["media"].Execute(w, nil)
}

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/public/img/"+r.URL.String()[len(`/images/`):])
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	ytemplate.ThePool.Fill("index", "templates/layout.html", "templates/index.html")

	ytemplate.ThePool.Pools["index"].Execute(w, nil)
}

func New() {
	wd, err := os.Getwd()
	yeasy.CheckError(err)

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/images/", ImageHandler)
	http.HandleFunc("/media", MediaHandler)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(wd+`/public`))))

	err = http.ListenAndServe(":9999", nil)
	yeasy.CheckError(err)
}
