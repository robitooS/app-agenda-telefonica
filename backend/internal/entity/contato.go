package entity

type Contato struct {
	ID        int64      `json:"id"`
	Nome      string     `json:"nome"`
	Idade     int        `json:"idade"`
	Telefones []Telefone `json:"telefones,omitempty"`
}
