package models

import "mime/multipart"

type RequestFile struct {
	Name       string                `json:"name"`
	Metadata   map[string]string     `json:"metadata,omitempty"`
	FileStream multipart.File        `json:"-"`
	FileHeader *multipart.FileHeader `json:"-"`
}
