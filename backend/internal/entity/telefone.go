package entity

type Telefone struct {
	IDContato int64  `json:"id_contato"`
	ID        int64  `json:"id"`
	Numero    string `json:"numero"`
}
