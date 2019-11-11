package main

import (
	"net/http"
	"sinosigorest/routes"
)

func main() {
	routes.CarregarRotas()
	http.ListenAndServe(":8000", nil)
}
