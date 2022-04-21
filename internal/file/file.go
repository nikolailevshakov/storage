package file

import (
	"github.com/google/uuid"
)

type File struct {
	id   uuid.UUID
	Name string
	Data []byte
}

func NewFile(filename string, blob []byte) (*File, error) {
	fileID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &File{
		id:   fileID,
		Name: filename,
		Data: blob,
	}, nil
}
