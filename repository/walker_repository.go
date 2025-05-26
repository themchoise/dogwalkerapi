package repository

import (
	"dogwalkerapi/model"
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"
)

type WalkerRepositoryI interface {
	Save(*model.JugadasData) error
	Read() ([]byte, error)
}

type WalerRepositoryImp struct {
	mu sync.RWMutex
}

func (r *WalerRepositoryImp) Read() (data []byte, err error) {

	r.mu.RLock()
	defer r.mu.RUnlock()

	readFile, err := os.Open("jugadas.json")

	if err != nil {
		log.Printf("Error abriendo archivo: %v", err)
		return nil, err
	}
	defer readFile.Close()

	bytes, err := io.ReadAll(readFile)
	if err != nil {
		log.Printf("Error leyendo archivo: %v", err)
		return nil, err
	}
	return bytes, nil

}

func (r *WalerRepositoryImp) Save(jugadasData *model.JugadasData) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	save, err := os.OpenFile("jugadas.json", os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		log.Printf("Error escriendo archivo para escritura: %v", err)
		return err
	}
	defer save.Close()

	encoder := json.NewEncoder(save)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(jugadasData); err != nil {
		log.Printf("Error escribiendo json: %v", err)
		return err
	}
	return nil

}

func NewWalkerRepository() WalkerRepositoryI {
	return &WalerRepositoryImp{}
}
