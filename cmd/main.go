package main

import (
	"Project/internal/settings"
	"Project/internal/settings/database"
	"Project/internal/settings/rout"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func init() {
	settings.LoadEnvVariables()
}

func main() {
	db, err := database.InitDB()

	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	r := chi.NewMux()
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	server := http.Server{Addr: port, Handler: r}

	log.Println("START")

	if err = rout.StartListener(&server, db, r); err != nil {
		log.Fatal(err.Error())
	}

}
