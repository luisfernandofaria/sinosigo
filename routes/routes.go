package routes

import (
	"net/http"
	"sinosigorest/controller"
)

func CarregarRotas() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new", controller.New)
	http.HandleFunc("/insert", controller.Insert)
}
