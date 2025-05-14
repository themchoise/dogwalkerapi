package routes

import (
	"dogwalkerapi/controller"
	"net/http"
)

type WalkerRouterImp struct {
	walkerController controller.WalkerControllerI
}

type WalkerRouterI interface {
	RegisterRoutes()
}

func NewWalkerRouter(ctrl controller.WalkerControllerI) WalkerRouterI {
	return &WalkerRouterImp{
		walkerController: ctrl,
	}
}

func (w *WalkerRouterImp) RegisterRoutes() {

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/walkers", w.walkerController.Hello)
	http.HandleFunc("/rungame", w.walkerController.RunGame)
}
