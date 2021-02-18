package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed frontend/public
var web embed.FS

//go:embed frontend/public/index.html
var index []byte

func main() {
	/* Add all static folders index.html is referencing to as subs */
	build, err := fs.Sub(web, "frontend/public/build")
	if err != nil {
		panic(err)
	}

	css, err := fs.Sub(web, "frontend/public/css")
	if err != nil {
		panic(err)
	}

	/* Single page application handler */
	http.HandleFunc("/", spaHandler)

	/* All static folder handlers */
	http.Handle("/build/", http.StripPrefix("/build", http.FileServer(http.FS(build))))
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.FS(css))))

	/* Additional routes like api fopr example */
	http.Handle("/api", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))

	/* Start server */
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func spaHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(index)
}
