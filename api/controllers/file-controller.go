package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/EdimarRibeiro/loadexcel/api/common"
	"github.com/EdimarRibeiro/loadexcel/internal/entities"
	entitiesinterface "github.com/EdimarRibeiro/loadexcel/internal/interfaces/entities"
)

type fileController struct {
	file entitiesinterface.FileRepositoryInterface
}

func CreateFileController(fileRep entitiesinterface.FileRepositoryInterface) *fileController {
	return &fileController{file: fileRep}
}

func (repo *fileController) GetAll(tenantId uint64) ([]entities.File, error) {
	files, err := repo.file.Search("")
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (repo *fileController) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	_, tenantId, err := common.ValidateToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	files, err := repo.GetAll(tenantId)
	if err != nil {
		http.Error(w, "Error retrieving file "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(files)
}

func (file *fileControler) CreateFileHandler(w http.ResponseWriter, r *http.Request) {
	var fileNew models.File
	if err := json.NewDecoder(r.Body).Decode(&fileNew); err != nil {
		http.Error(w, "Error decoding JSON "+err.Error(), http.StatusBadRequest)
		return
	}

	err := file.CreateNewFile(fileNew)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
