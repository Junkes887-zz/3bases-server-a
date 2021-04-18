package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UsuarioDecrypt struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CPF         string             `json:"cpf"`
	Nome        string             `json:"nome"`
	Endereco    string             `json:"endereco"`
	ListDividas []DividaDecrypt    `json:"listDividas"`
}

type DividaDecrypt struct {
	Descricao string `json:"descricao"`
}

type UsuarioEncrypt struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CPF         []byte             `bson:"cpf"`
	Nome        []byte             `bson:"nome"`
	Endereco    []byte             `bson:"endereco"`
	ListDividas []DividaEncrypt    `bson:"listDividas"`
}

type DividaEncrypt struct {
	Descricao []byte `bson:"descricao"`
}
