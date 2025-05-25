package model

type JugadasData struct {
	Piedra int `json:"piedra"`
	Papel  int `json:"papel"`
	Tijera int `json:"tijera"`
}

func (j *JugadasData) JugadasDataBetter() string {
	if j.Papel > j.Piedra && j.Papel > j.Tijera {
		return "Tijera"
	}
	if j.Piedra > j.Papel && j.Piedra > j.Tijera {
		return "Papel"
	}
	return "Piedra"

}
