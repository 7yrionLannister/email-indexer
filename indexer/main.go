package main

import (
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := configureChiRouter()
	log.Println(http.ListenAndServe("localhost:6060", r))
}

func configureChiRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/process-emails", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		// iterate all files and directories in the user folder recursively and process the files
		processEmailFile("maildir/" + name)
		w.Write([]byte("Finished processing emails!"))
	})
	r.Get("/get-emails", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // allow the Vue application to access this method
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(fetchEmails(name)))
	})
	r.HandleFunc("/debug/heap", func(w http.ResponseWriter, r *http.Request) {
		pprof.Handler("heap").ServeHTTP(w, r)
	})
	r.HandleFunc("/debug/cpu", func(w http.ResponseWriter, r *http.Request) {
		pprof.Profile(w, r)
	})
	return r
}
