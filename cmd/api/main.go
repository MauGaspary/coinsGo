package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/MauGaspary/goapi/internal/handlers"
	"github.com/MauGaspary/goapi/internal/tools"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	var err error

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		log.Fatal(err)
		return
	}

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r, *database)
	
	// Servir index.html na raiz
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})
	
	// Servir arquivos estáticos
	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static", fs))
	// Também servir CSS e JS na raiz para compatibilidade
	r.Get("/{file}", func(w http.ResponseWriter, r *http.Request) {
		file := chi.URLParam(r, "file")
		http.ServeFile(w, r, "./static/"+file)
	})
	
	fmt.Println("Starting API server on port", port)
	fmt.Println(`
	 ██████   ██████       █████  ██████  ██ 
	██       ██    ██     ██   ██ ██   ██ ██ 
	██   ███ ██    ██     ███████ ██████  ██ 
	██    ██ ██    ██     ██   ██ ██      ██ 
	 ██████   ██████      ██   ██ ██      ██                                      
	`)

	fmt.Println("Server is now running.")

	err = http.ListenAndServe(":" + port, r)
	if err != nil {
		log.Error(err)
	}
}