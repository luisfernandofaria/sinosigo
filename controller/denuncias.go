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

		latitude := r.FormValue("latitude")
		longitude := r.FormValue("longitude")
		endereco := r.FormValue("endereco")
		municipio := r.FormValue("municipio")
		cep := r.FormValue("cep")

		var localacidente = model.CriaNovoLocalAcidente(latitude, longitude, endereco, municipio, cep)

		descricao := r.FormValue("descricao")
		data := r.FormValue("data")
		imagem := r.FormValue("imagem")
		autordodano := r.FormValue("autordodano")
		emailusuario := r.FormValue("emailusuario")
		categoria := r.FormValue("categoria")

		model.CriaNovaDenuncia(descricao, data, imagem, autordodano, emailusuario, categoria, localacidente)
	}
	http.Redirect(w, r, "/", 301)
}
