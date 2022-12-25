package main

import (
	"log"
	"net/http"

	"github.com/mrcn04/godo/pkg/api"
	"github.com/mrcn04/godo/pkg/db"
	"github.com/mrcn04/godo/pkg/handlers"
)

func main() {
	init := db.InitDatabase("")

	defer init.DB.Close()

	h := handlers.NewHandler(init.DB)
	r := api.RegisterRoutes(h)

	log.Println("Server started on " + init.Port)
	log.Fatal(http.ListenAndServe(":"+init.Port, r))
}
