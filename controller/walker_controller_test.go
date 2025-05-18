package controller

import "testing"

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
