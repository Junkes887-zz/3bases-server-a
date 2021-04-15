package builder

import (
	artifacts "github.com/Junkes887/3bases-artifacts"
	"github.com/Junkes887/3bases-server-a/model"
)

func EncryptUsuario(usuario model.UsuarioDecrypt) model.UsuarioEncrypt {
	return model.UsuarioEncrypt{
		CPF:         artifacts.Encrypt(usuario.CPF, usuario.CPF),
		Nome:        artifacts.Encrypt(usuario.Nome, usuario.CPF),
		Endereco:    usuario.Endereco,
		ListDividas: usuario.ListDividas,
	}
}

func DecryptUsuario(usuario model.UsuarioEncrypt, cpf string) model.UsuarioDecrypt {
	return model.UsuarioDecrypt{
		CPF:         artifacts.Decrypt(usuario.CPF, cpf),
		Nome:        artifacts.Decrypt(usuario.Nome, cpf),
		Endereco:    usuario.Endereco,
		ListDividas: usuario.ListDividas,
	}
}
