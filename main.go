package main

import (
	"fmt"
	"net/http"

	"github.com/EdimarRibeiro/loadexcel/api/routers"
	"github.com/EdimarRibeiro/loadexcel/internal/infrastructure/database"
	"github.com/gorilla/mux"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Name, cache-control, postman-token")
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	database.Initialize(false)
	router := mux.NewRouter()

	// Aplicar el middleware CORS a todas las rutas
	router.Use(corsMiddleware)

	/*************************************/
	routers.CreateRouterOthers(router, database.DB)

	port := 8085
	fmt.Printf("Server started on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)

}
