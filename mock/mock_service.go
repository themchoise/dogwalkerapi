package mock

import (
	"dogwalkerapi/model"
)

type MockWalkerService struct {
	WriteFileFunc func(*model.JugadasData) error
	OpenFileFunc  func() ([]byte, error)
}

func (m *MockWalkerService) WriteFile(j *model.JugadasData) error {
	if m.WriteFileFunc != nil {
		return m.WriteFileFunc(j)
	}
	return nil
}

func (m *MockWalkerService) OpenFile() ([]byte, error) {
	if m.OpenFileFunc != nil {
		return m.OpenFileFunc()
	}
	return nil, nil
}


