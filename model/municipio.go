package model

type Municipio struct {
	ID         int64 `json:"id"`
	Descricao  string `json:"descricao"`
	Uf         string `json:"uf"`
	CodigoIbge int64 `json:"codigoibge"`
}
