package database

import (
	"github.com/EdimarRibeiro/loadexcel/internal/entities"
	"gorm.io/gorm"
)

type FileRepository struct {
	DB *gorm.DB
}

func CreateFileRepository(db *gorm.DB) *FileRepository {
	return &FileRepository{DB: db}
}

func (entity *FileRepository) Save(model *entities.File) (*entities.File, error) {
	result := entity.DB.Save(&model)
	return model, result.Error
}

func (entity *FileRepository) GetFileId(value string) (uint64, error) {
	var model entities.File
	result := entity.DB.First(&model, "Id = ?", value)
	if result.Error != nil {
		return 0, result.Error
	}
	return model.Id, nil
}

func (entity *FileRepository) Search(where string) ([]entities.File, error) {
	var model []entities.File
	result := entity.DB.Where(where).Find(&model)
	if result.Error != nil {
		return nil, result.Error
	}
	return model, nil
}
