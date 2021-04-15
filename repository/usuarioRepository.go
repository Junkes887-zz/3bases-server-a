package repository

import (
	"context"
	"fmt"

	"github.com/Junkes887/3bases/server-base-a/crypt"
	"github.com/Junkes887/3bases/server-base-a/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	DB  *mongo.Collection
	CTX context.Context
}

func (client Client) Save(usuario model.UsuarioDecrypt) {
	usuarioEncrypt := crypt.EncryptUsuario(usuario)

	_, err := client.DB.InsertOne(client.CTX, usuarioEncrypt)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (client Client) Find(cpf string) model.UsuarioDecrypt {
	var usuarios []model.UsuarioEncrypt
	cur, err := client.DB.Find(client.CTX, bson.D{})

	if err != nil {
		fmt.Print(err)
	}

	for cur.Next(client.CTX) {
		var u model.UsuarioEncrypt
		err := cur.Decode(&u)
		if err != nil {
			fmt.Println(err)
		}
		usuarios = append(usuarios, u)
	}

	if err := cur.Err(); err != nil {
		fmt.Println(err)
	}

	cur.Close(client.CTX)

	for _, usuario := range usuarios {
		usuarioDecrypt := crypt.DecryptUsuario(usuario, cpf)
		if usuarioDecrypt.CPF != "" {
			return usuarioDecrypt
		}
	}

	return model.UsuarioDecrypt{}
}
