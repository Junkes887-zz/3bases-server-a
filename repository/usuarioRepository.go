package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/Junkes887/3bases-server-a/builder"
	"github.com/Junkes887/3bases-server-a/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	DB  *mongo.Collection
	CTX context.Context
}

func (client Client) FindAll() []model.UsuarioDecrypt {
	var usuarios []model.UsuarioEncrypt
	var usuariosDecrypt []model.UsuarioDecrypt
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
		usuarioDecrypt := builder.DecryptUsuario(usuario)
		usuariosDecrypt = append(usuariosDecrypt, usuarioDecrypt)
	}

	return usuariosDecrypt
}

func (client Client) Find(id string) model.UsuarioDecrypt {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid id")
	}

	var usuario model.UsuarioEncrypt
	filter := bson.M{
		"_id": objectId,
	}
	err = client.DB.FindOne(client.CTX, filter).Decode(&usuario)
	if err != nil {
		log.Println(err)
	}

	if usuario.ID == primitive.NilObjectID {
		return model.UsuarioDecrypt{}
	}

	usuarioDecrypt := builder.DecryptUsuario(usuario)
	return usuarioDecrypt
}

func (client Client) Save(usuario model.UsuarioDecrypt) interface{} {
	usuarioEncrypt := builder.EncryptUsuario(usuario)

	res, err := client.DB.InsertOne(client.CTX, usuarioEncrypt)
	if err != nil {
		fmt.Println(err)
	}

	return res.InsertedID
}

func (client Client) Upadate(id string, usuario model.UsuarioDecrypt) string {
	usuarioEncrypt := builder.EncryptUsuario(usuario)

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println("Invalid id")
	}

	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "cpf", Value: usuarioEncrypt.CPF},
		primitive.E{Key: "nome", Value: usuarioEncrypt.Nome},
		primitive.E{Key: "endereco", Value: usuarioEncrypt.Endereco},
		primitive.E{Key: "listDividas", Value: usuarioEncrypt.ListDividas},
	}}}

	res, err := client.DB.UpdateByID(client.CTX, objectId, update)

	if err != nil {
		log.Println("Invalid id")
	}

	if res.ModifiedCount == 0 {
		return "Usuario não encontrado"
	}

	return "Usuario atualizado"
}

func (cliente Client) Delete(id string) string {
	objectId, err := primitive.ObjectIDFromHex(id)
	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}

	res, err := cliente.DB.DeleteOne(cliente.CTX, filter)
	if err != nil {
		log.Println(err)
	}

	if res.DeletedCount == 0 {
		return "Usuario não encontrado"
	}

	return "Usuario removido"
}
