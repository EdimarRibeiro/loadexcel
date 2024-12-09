package entities

import (
	"errors"
	"io"
)

/**/
type File struct {
	Id       uint64            `gorm:"primaryKey;autoIncrement:true"  json:"id"`
	Name     string            `gorm:"size:150" json:"name"`
	Metadata map[string]string `gorm:"-" json:"metadata,omitempty"`
	// Stream do arquivo (não persistido no banco de dados)
	FileStream io.Reader `gorm:"-" json:"-"`
}

func (c *File) Validate() error {
	if c.Name == "" {
		return errors.New("the name is required")
	}
	return nil
}

func NewFile(name string, metadata map[string]string, fileStream io.Reader) (*File, error) {
	file := &File{
		Name:       name,
		Metadata:   metadata,
		FileStream: fileStream,
	}

	// Validação do arquivo
	if err := file.Validate(); err != nil {
		return nil, err
	}

	return file, nil
}
