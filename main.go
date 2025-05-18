package main

import (
	"dogwalkerapi/controller"
	"dogwalkerapi/routes"
	"log"
	"net/http"
)

func main() {

	walkerController := controller.NewWalkerController()
	walkerRouter := routes.NewWalkerRouter(walkerController)
	walkerRouter.RegisterRoutes()

	port := ":8080"
	log.Printf("Servidor escuchando en http://localhost%s", port)

	http.Handle("/", http.RedirectHandler("/hello", 302))

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
