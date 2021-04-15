package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Junkes887/3bases-server-a/model"
	"github.com/Junkes887/3bases-server-a/repository"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	DB  *mongo.Collection
	CTX context.Context
	REP repository.Client
}

func (client Client) Find(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cpf := p.ByName("cpf")

	repository := client.REP
	usuario := repository.Find(cpf)

	js, err := json.Marshal(usuario)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (client Client) Save(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var usuario model.UsuarioDecrypt

	json.NewDecoder(r.Body).Decode(&usuario)
	client.REP.Save(usuario)

	js, err := json.Marshal("Usuario criado!!!")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(js)
}
