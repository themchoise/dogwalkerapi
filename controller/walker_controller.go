package controller

import (
	"dogwalkerapi/config"
	"dogwalkerapi/model"
	"dogwalkerapi/service"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"slices"
)

type WalkerControllerImp struct {
	Serv service.WalkerServiceI
}

const PIEDRA = "Piedra"
const PAPEL = "Papel"
const TIJERA = "Tijera"

type PageData struct {
	Title   string
	Message string
}

type ResultadoDTO struct {
	Resultado       string `json:"resultado"`
	IsPlayerVictory bool   `json:"isPlayerVictory"`
	JugadaPC        string `json:"jugadaPC"`
}

func IsPlayerVictory(player string, computer string) bool {

	if player == computer {
		return false
	}

	switch player {

	case PIEDRA:
		if computer == PAPEL {
			return false
		}
	case PAPEL:
		if computer == TIJERA {
			return false
		}
	case TIJERA:
		if computer == PIEDRA {
			return false
		}
	}

	return true

}

func GetTemplateConfigFromConfig() config.TemplateConfigI {
	return config.GetTemplateConfig()
}

type WalkerControllerI interface {
	Hello(w http.ResponseWriter, r *http.Request)
	RunGame(w http.ResponseWriter, r *http.Request)
	Play(w http.ResponseWriter, r *http.Request)
}

func NewWalkerController(serv service.WalkerServiceI) WalkerControllerI {
	return &WalkerControllerImp{
		Serv: serv,
	}
}

func (*WalkerControllerImp) Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	data := PageData{
		Title:   "Mi Página Web",
		Message: "¡Hola desde Go Templates!",
	}

	tmpl, err := template.ParseFiles("./templates/layout.html", "./templates/index.html")
	if err != nil {
		log.Fatalf("Error al cargar plantillas: %v", err)
	}

	err = tmpl.ExecuteTemplate(w, "layout", GetTemplateConfigFromConfig().Config())
	if err != nil {
		log.Fatal("Error al ejecutar la plantilla: ", err)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (*WalkerControllerImp) RunGame(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	tmpl, err := template.ParseFiles("./templates/layout.html", "./templates/rungame.html")
	if err != nil {
		log.Fatalf("Error al cargar plantillas: %v", err)
	}

	err = tmpl.ExecuteTemplate(w, "layout", GetTemplateConfigFromConfig().Config())
	if err != nil {
		log.Fatal("Error al ejecutar la plantilla: ", err)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func (c *WalkerControllerImp) Play(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var jugadasData model.JugadasData
	options := []string{"Piedra", "Papel", "Tijera"}
	playerOption := r.Header.Get("jugada")

	isValidOption := slices.Contains(options, playerOption)

	if !isValidOption {
		http.Error(w, "Opción de Jugada inválida", http.StatusBadRequest)
		return
	}

	bytes, err := c.Serv.OpenFile()
	if err != nil {
		http.Error(w, "Error interno de servidor", http.StatusInternalServerError)
		return
	}

	if len(bytes) > 0 {
		if err := json.Unmarshal(bytes, &jugadasData); err != nil {
			log.Fatal("Error parseando JSON:", err)
		}
	}

	switch playerOption {
	case "Piedra":
		jugadasData.Piedra++
	case "Papel":
		jugadasData.Papel++
	case "Tijera":
		jugadasData.Tijera++
	}

	computerOption := (&jugadasData).JugadasDataBetter()

	err = c.Serv.WriteFile(&jugadasData)
	if err != nil {
		log.Fatal("Error al Escribir Archivo")
	}

	resultado := ResultadoDTO{
		Resultado:       "La Pc Jugo: " + computerOption,
		IsPlayerVictory: IsPlayerVictory(playerOption, computerOption),
		JugadaPC:        computerOption,
	}

	json.NewEncoder(w).Encode(resultado)
}
