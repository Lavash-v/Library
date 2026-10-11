package storage

import (
	"encoding/json"
	"library-app/domain"
	"library-app/library"
	"os"
)

type Storable interface {
	Save() error
	Load() error
}

type storageData struct {
	Books []*domain.Book`json:"books"`
	Readers []*domain.Reader`json:"readers"`
}

func SaveLibraryToJSON(filePath string, lib *library.Library) error {
	data := storageData {
		Books: lib.Books,
		Readers: lib.Readers,
	}

	jsonData, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, jsonData, 0644)
}

func LoadLibraryFromJSON(filePath string) (*library.Library, error) {
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var data storageData

	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, err
	}

	lib := &library.Library{
		Books:   data.Books,
		Readers: data.Readers,
	}
	return lib, nil
}

