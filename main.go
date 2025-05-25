package main

import (
	"dogwalkerapi/controller"
	"dogwalkerapi/repository"
	"dogwalkerapi/routes"
	"dogwalkerapi/service"
	"log"
	"net/http"
)

func main() {

	repo := repository.NewWalkerRepository()
	service := service.NewWalkerService(repo)
	walkerController := controller.NewWalkerController(service)
	walkerRouter := routes.NewWalkerRouter(walkerController)
	walkerRouter.RegisterRoutes()

	port := ":8080"
	log.Printf("Servidor escuchando en http://localhost%s", port)

	http.Handle("/", http.RedirectHandler("/hello", 302))

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
