package model

import "sinosigorest/database"

type Denuncia struct {
	ID            int64 `json:"id"`
	LocalAcidente *LocalAcidente `json:"localacidente"`
	Descricao     string `json:"descricao"`
	Data          string`json:"data"`
	Foto          string`json:"foto"`
	AutorDano     string`json:"autordano"`
	EmailUsuario  string`json:"emailusuario"`
	Categoria     string`json:"categoria"`
}

func BuscarDenuncias() []Denuncia {
	db := database.ConectarComBanco()

	selectDenuncias, err:= db.Query("SELECT id, descricao, datadenuncia, categoria FROM denuncia")
	if err != nil {
	panic(err.Error())
	}

 	d:= Denuncia{}
	denuncias := []Denuncia{}

	for selectDenuncias.Next() {
	var id int64
		var descricao, data, categoria string

		err = selectDenuncias.Scan(&id, &descricao, &data, &categoria)
		if err != nil {
		panic(err.Error())
		}

		d.Descricao = descricao
		d.Data = data
		d.Categoria = categoria
		denuncias = append(denuncias, d)
	}
	defer db.Close()
	return denuncias

}

func CriaNovaDenuncia(descricao string, data string, imagem string, autordodano string, emailusuario string, categoria string, localacidente LocalAcidente) {

	db := database.ConectarComBanco()

	insereDadosNoBanco, err := db.Prepare("insert into denuncia(descricao, datadenuncia, imagem, autordodano, emailusuario, categoria, localacidente_denuncia_id) values($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(descricao, data, imagem, autordodano, emailusuario, categoria, localacidente)
	defer db.Close()
}
