package controller

import (
	"net/http"
	"text/template"

	"sinosigorest/model"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	todosOsProdutos := model.BuscarDenuncias()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		descricao := r.FormValue("descricao")
		data := r.FormValue("data")
		imagem := r.FormValue("imagem")
		autordodano := r.FormValue("autordodano")
		emailusuario := r.FormValue("emailusuario")
		categoria := r.FormValue("categoria")
		LocalAcidente := r.FormValue("localacidente")

		model.CriaNovaDenuncia(descricao, data, imagem, autordodano, emailusuario, categoria, LocalAcidente)
	}
	http.Redirect(w, r, "/", 301)
}