package controller

import (
	"dogwalkerapi/config"
	"html/template"
	"log"
	"net/http"
)

type WalkerControllerImp struct{}

type PageData struct {
	Title   string
	Message string
}

func GetTemplateConfigFromConfig() config.TemplateConfigI {
	return config.GetTemplateConfig()
}

type WalkerControllerI interface {
	Hello(w http.ResponseWriter, r *http.Request)
	RunGame(w http.ResponseWriter, r *http.Request)
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
		return
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
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
