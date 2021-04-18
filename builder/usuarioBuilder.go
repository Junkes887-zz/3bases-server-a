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
		Endereco:    artifacts.Encrypt(usuario.Endereco),
		ListDividas: encryptListDividas(usuario.ListDividas),
	}
}

func encryptListDividas(dividas []model.DividaDecrypt) []model.DividaEncrypt {
	var divadasEncrypt []model.DividaEncrypt
	for _, divida := range dividas {
		divadaEncrypt := model.DividaEncrypt{
			Descricao: artifacts.Encrypt(divida.Descricao),
		}
		divadasEncrypt = append(divadasEncrypt, divadaEncrypt)
	}

	return divadasEncrypt
}

func DecryptUsuario(usuario model.UsuarioEncrypt) model.UsuarioDecrypt {
	return model.UsuarioDecrypt{
		ID:          usuario.ID,
		CPF:         artifacts.Decrypt(usuario.CPF),
		Nome:        artifacts.Decrypt(usuario.Nome),
		Endereco:    artifacts.Decrypt(usuario.Endereco),
		ListDividas: DecryptListDividas(usuario.ListDividas),
	}
}

func DecryptListDividas(dividas []model.DividaEncrypt) []model.DividaDecrypt {
	var divadasDecrypt []model.DividaDecrypt
	for _, divida := range dividas {
		divadaDecrypt := model.DividaDecrypt{
			Descricao: artifacts.Decrypt(divida.Descricao),
		}
		divadasDecrypt = append(divadasDecrypt, divadaDecrypt)
	}

	return divadasDecrypt
}
