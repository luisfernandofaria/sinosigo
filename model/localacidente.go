package model

import (
	"sinosigorest/database"
)

type LocalAcidente struct {
	ID        int64 `json:"id"`
	Latitude  string`json:"latitude"`
	Longitude string`json:"longitude"`
	Endereco  string`json:"endereco"`
	Municipio *Municipio `json:"municipio"`
	Cep       string`json:"cep"`
}

func CriaNovoLocalAcidente(latitude string, longitude string, endereco string, municipio Municipio, cep string) {
	db := database.ConectarComBanco()

	insereDadosNoBanco, err := db.Prepare("insert into localacidente(latitude, longitude, endereco, municipio, cep) values($1, $2, $3, $4, $5)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(latitude, longitude, endereco, municipio, cep)
	defer db.Close()
}
