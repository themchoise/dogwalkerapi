package service

import (
	"dogwalkerapi/model"
	"dogwalkerapi/repository"
)

type WalkerServiceI interface {
	WriteFile(*model.JugadasData) error
	OpenFile() ([]byte, error)
}

type WalkerService struct {
	Repo repository.WalkerRepositoryI
}

func (w *WalkerService) OpenFile() ([]byte, error) {
	return w.Repo.Read()
}

func (w *WalkerService) WriteFile(jugada *model.JugadasData) error {
	return w.Repo.Save(jugada)
}

func NewWalkerService(repo repository.WalkerRepositoryI) WalkerServiceI {
	return &WalkerService{
		Repo: repo,
	}
}
