package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UsuarioDecrypt struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CPF         string             `json:"cpf"`
	Nome        string             `json:"nome"`
	Endereco    string             `json:"endereco"`
	ListDividas []string           `json:"listDividas"`
}

type UsuarioEncrypt struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CPF         []byte             `bson:"cpf"`
	Nome        []byte             `bson:"nome"`
	Endereco    string             `bson:"endereco"`
	ListDividas []string           `bson:"listDividas"`
}
