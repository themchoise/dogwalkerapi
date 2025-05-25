package repository

import (
	"dogwalkerapi/model"
	"encoding/json"
	"io"
	"log"
	"os"
)

type WalkerRepositoryI interface {
	Save(*model.JugadasData) error
	Read() ([]byte, error)
}

type WalkerControllerImp struct{}

func (w *WalkerControllerImp) Read() (data []byte, err error) {

	log.Printf("Hola")

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

func (w *WalkerControllerImp) Save(jugadasData *model.JugadasData) error {

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
	return &WalkerControllerImp{}
}
