package builder

import (
	artifacts "github.com/Junkes887/3bases-artifacts"
	"github.com/Junkes887/3bases-server-a/model"
)

func EncryptUsuario(usuario model.UsuarioDecrypt) model.UsuarioEncrypt {
	return model.UsuarioEncrypt{
		ID:          usuario.ID,
		CPF:         artifacts.Encrypt(usuario.CPF),
		Nome:        artifacts.Encrypt(usuario.Nome),
		Endereco:    usuario.Endereco,
		ListDividas: usuario.ListDividas,
	}
}

func DecryptUsuario(usuario model.UsuarioEncrypt) model.UsuarioDecrypt {
	return model.UsuarioDecrypt{
		ID:          usuario.ID,
		CPF:         artifacts.Decrypt(usuario.CPF),
		Nome:        artifacts.Decrypt(usuario.Nome),
		Endereco:    usuario.Endereco,
		ListDividas: usuario.ListDividas,
	}
}
