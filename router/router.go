package router

import (
	"barber/router/routes"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Generate() {
	r := mux.NewRouter()
	routes.RoutesConfiguration(r)
	log.Fatal(
		http.ListenAndServe(":8000", handlers.CORS(
			handlers.AllowedOrigins(
				[]string{"http://localhost:3333"},
			),
		)(r)),
	)
}
