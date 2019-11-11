package model

import "sinosigorest/database"

type Denuncia struct {
	ID            int64
	LocalAcidente string
	Descricao     string
	Data          string
	Foto          string
	AutorDano     string
	EmailUsuario  string
	Categoria     string
}

func BuscarDenuncias() []Denuncia {
	db := database.ConectarComBanco()

	selectDenuncias, err := db.Query("SELECT id, descricao, datadenuncia, categoria FROM denuncia")
	if err != nil {
		panic(err.Error())
	}

	d := Denuncia{}
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

func CriaNovaDenuncia(descricao string, data string, categoria string) {
	db := database.ConectarComBanco()

	insereDadosNoBanco, err := db.Prepare("insert into denuncia(descricao, datadenuncia, categoria) values($1, $2, $3)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(descricao, data, categoria)
	defer db.Close()
}
