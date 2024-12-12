package routers

import (
	"github.com/EdimarRibeiro/loadexcel/api/controllers"
	"github.com/gorilla/mux"
)

// CreateRouterFile configura rotas relacionadas a arquivos
func CreateRouterFile(router *mux.Router) {
	fileController := controllers.CreateFileController()

	router.HandleFunc("/api/file", fileController.CreateFileHandler).Methods("OPTIONS", "POST")
}
