package controller

import (
	"dogwalkerapi/config"
	"html/template"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func TestJugadasDataBetter(t *testing.T) {
	tests := []struct {
		nombre   string
		data     JugadasData
		esperado string
	}{
		{"Predice Tijera", JugadasData{Piedra: 1, Papel: 2, Tijera: 3}, "Piedra"},
		{"Predice Tijera", JugadasData{Piedra: 2, Papel: 3, Tijera: 2}, "Tijera"},
		{"Predice Papel", JugadasData{Piedra: 3, Papel: 2, Tijera: 1}, "Papel"},
	}

	for _, tt := range tests {
		t.Run(tt.nombre, func(t *testing.T) {
			resultado := tt.data.JugadasDataBetter()
			if resultado != tt.esperado {
				t.Errorf("esperado %s, obtuviste %s", tt.esperado, resultado)
			}
		})
	}
}

func TestIsPlayerVictory(t *testing.T) {
	tests := []struct {
		testname string
		player   string
		computer string
		esperado bool
	}{
		{"Jugador juega " + PIEDRA + " y la computadora " + PIEDRA, PIEDRA, PIEDRA, false},
		{"Jugador juega " + PIEDRA + " y la computadora " + PAPEL, PIEDRA, PAPEL, false},
		{"Jugador juega " + PIEDRA + " y la computadora " + TIJERA, PIEDRA, TIJERA, true},

		{"Jugador juega " + PAPEL + " y la computadora " + PAPEL, PAPEL, PAPEL, false},
		{"Jugador juega " + PAPEL + " y la computadora " + TIJERA, PAPEL, TIJERA, false},
		{"Jugador juega " + PAPEL + " y la computadora " + PIEDRA, PAPEL, PIEDRA, true},

		{"Jugador juega " + TIJERA + " y la computadora " + PAPEL, TIJERA, PAPEL, true},
		{"Jugador juega " + TIJERA + " y la computadora " + TIJERA, TIJERA, TIJERA, false},
		{"Jugador juega " + TIJERA + " y la computadora " + PIEDRA, TIJERA, PIEDRA, false},
	}

	for _, tt := range tests {
		t.Run(tt.testname, func(t *testing.T) {
			resultado := IsPlayerVictory(tt.player, tt.computer)

			if resultado != tt.esperado {
				t.Errorf("Se esperaba %v, obtuviste %v", tt.esperado, resultado)
			}
		})
	}
}

func TestGetTemplateConfigFromConfig(t *testing.T) {

	expectedType := reflect.TypeOf(&config.TemplateConfigImp{})

	resultType := reflect.TypeOf(GetTemplateConfigFromConfig())

	if resultType != expectedType {
		t.Errorf("La estructura del resultado no coincide. Esperado: %v, Obtenido: %v", expectedType, resultType)
	}

}

func TestNewWalkerController(t *testing.T) {

	expectedType := reflect.TypeOf(&WalkerControllerImp{})

	resultType := reflect.TypeOf(NewWalkerController())

	if resultType != expectedType {
		t.Errorf("La estructura del resultado no coincide. Esperado: %v, Obtenido: %v", expectedType, resultType)
	}

}

func TestHello(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title:   "Mi Página Web",
			Message: "¡Hola desde Go Templates!",
		}

		_, filename, _, _ := runtime.Caller(0)
		base := filepath.Join(filepath.Dir(filename), "..", "templates")

		layout := filepath.Join(base, "layout.html")
		index := filepath.Join(base, "index.html")

		tmpl, err := template.ParseFiles(layout, index)
		if err != nil {
			http.Error(w, "Error interno al renderizar la plantilla", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, data)
	})

	handler.ServeHTTP(rr, req)


	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code incorrecto. Esperado: %d, Obtenido: %d", http.StatusOK, status)
	}

	body := rr.Body.String()
	if !strings.Contains(body, "") {
		t.Errorf("No se encontró el título esperado en el HTML")
	}
	if !strings.Contains(body, "") {
		t.Errorf("No se encontró el mensaje esperado en el HTML")
	}
}
