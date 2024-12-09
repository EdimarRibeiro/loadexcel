package routers

import (
	"github.com/EdimarRibeiro/loadexcel/api/controllers"
	"github.com/EdimarRibeiro/loadexcel/internal/infrastructure/database"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CreateRouterFile(router *mux.Router, db *gorm.DB) {
	fileRepo := database.CreateFileRepository(db)
	file := controllers.CreateFileController(fileRepo)

	// Private route to get all users (requires JWT)
	router.HandleFunc("/api/files", file.GetAllHandler).Methods("OPTIONS")
	router.HandleFunc("/api/files", file.GetAllHandler).Methods("GET")

	// Handle POST requests to /api/file with the CreateFileHandler function
	router.HandleFunc("/api/file", file.CreateFileHandler).Methods("OPTIONS")
	router.HandleFunc("/api/file", file.CreateFileHandler).Methods("POST")
}
