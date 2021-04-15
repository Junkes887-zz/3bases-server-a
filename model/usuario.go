package model

type UsuarioDecrypt struct {
	CPF         string   `json:"cpf"`
	Nome        string   `json:"nome"`
	Endereco    string   `json:"endereco"`
	ListDividas []string `json:"listDividas"`
}

type UsuarioEncrypt struct {
	CPF         []byte   `bson:"cpf"`
	Nome        []byte   `bson:"nome"`
	Endereco    string   `bson:"endereco"`
	ListDividas []string `bson:"listDividas"`
}
