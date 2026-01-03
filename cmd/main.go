package main

import (
	"fmt"
	"net/http"

	"github.com/claytonCharles/albionatlas-api/database"
	"github.com/claytonCharles/albionatlas-api/internal/auth"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Alerta: Não foi possivel carregar as variaveis do sistema")
	}

	database := database.NewConnection()
	defer database.DB.Close()

	database.InitializeRepositories()
	authHandler := auth.NewHandler(database.AuthRepo)

	http.HandleFunc("POST /register", authHandler.Register)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	fmt.Println("Servidor rodando na porta 5656")
	http.ListenAndServe(":5656", nil)
}
