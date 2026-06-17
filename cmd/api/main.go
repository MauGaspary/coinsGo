package main

import (
	"fmt"
	"net/http"
	"os"
	"database/sql"

	"github.com/MauGaspary/goapi/internal/database"
	"github.com/MauGaspary/goapi/internal/handlers"
	// "github.com/MauGaspary/goapi/internal/tools"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)


// @title Projeto GoBank API
// @version 1.0
// @description API de serviços financeiros simulados
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("Não foi possível encontrar a variável de ambiente DATABASE_URL")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	dbQueries := database.New(dbConn)

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r, dbQueries)
	
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

	fmt.Println("Server is now running." + " Access the API at http://localhost:" + port)

	err = http.ListenAndServe(":" + port, r)
	if err != nil {
		log.Error(err)
	}
}