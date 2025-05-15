package controller

import (
	"dogwalkerapi/config"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"slices"
)

type WalkerControllerImp struct{}

type PageData struct {
	Title   string
	Message string
}

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

func GetTemplateConfigFromConfig() config.TemplateConfigI {
	return config.GetTemplateConfig()
}

type WalkerControllerI interface {
	Hello(w http.ResponseWriter, r *http.Request)
	RunGame(w http.ResponseWriter, r *http.Request)
	Play(w http.ResponseWriter, r *http.Request)
}

func NewWalkerController() WalkerControllerI {
	return &WalkerControllerImp{}
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
func (*WalkerControllerImp) Play(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	options := []string{"Piedra", "Papel", "Tijera"}
	optionSelected := r.Header.Get("jugada")

	isValidOption := slices.Contains(options, optionSelected)

	if !isValidOption {
		http.Error(w, "Opción inválida", http.StatusBadRequest)
		return
	}

	readFile, err := os.Open("jugadas.json")
	if err != nil {
		log.Fatal("Error abriendo archivo:", err)
	}
	defer readFile.Close()

	bytes, err := io.ReadAll(readFile)
	if err != nil {
		log.Fatal("Error leyendo archivo:", err)
	}

	var jugadasData JugadasData
	if len(bytes) > 0 {
		if err := json.Unmarshal(bytes, &jugadasData); err != nil {
			log.Fatal("Error parseando JSON:", err)
		}
	}

	switch optionSelected {
	case "Piedra":
		jugadasData.Piedra++
	case "Papel":
		jugadasData.Papel++
	case "Tijera":
		jugadasData.Tijera++
	}

	computerOption := (&jugadasData).JugadasDataBetter()

	writeFile, err := os.OpenFile("jugadas.json", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("Error abriendo archivo para escritura:", err)
	}
	defer writeFile.Close()

	encoder := json.NewEncoder(writeFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(jugadasData); err != nil {
		log.Fatal("Error escribiendo JSON:", err)
	}

	json.NewEncoder(w).Encode(map[string]string{
		"resultado": "La Pc Jugo: " + computerOption,
	})
}
