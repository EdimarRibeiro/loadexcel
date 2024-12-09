package entitiesinterface

import (
	entities "github.com/EdimarRibeiro/loadexcel/internal/entities"
)

type FileRepositoryInterface interface {
	Save(model *entities.File) (*entities.File, error)
	GetFileId(value string) (uint64, error)
	Search(where string) ([]entities.File, error)
}
