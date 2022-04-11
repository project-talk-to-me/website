package main

import (
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", getFile("./flyer/index.html"))
	r.Get("/index.html", getFile("./html/index.html"))
	r.Get("/*", getDir("./html/"))

	http.ListenAndServe(":8080", r)
}

func getFile(fn string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reader, err := os.Open(fn)
		if err != nil {
			panic(err)
		}
		io.Copy(w, reader)
	}
}

func getDir(path string) http.HandlerFunc {
	return http.FileServer(http.Dir(path)).ServeHTTP
}
