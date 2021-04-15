package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"os"

	"github.com/Junkes887/3bases/server-base-a/model"
)

func EncryptCpf(cpf string) string {
	CRYPT_KEY := os.Getenv("CRYPT_KEY")
	key := []byte(cpf + CRYPT_KEY)
	plaintext := []byte(cpf)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, aesgcm.NonceSize())
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	return string(ciphertext)
}

func EncryptUsuario(usuario model.UsuarioDecrypt) model.UsuarioEncrypt {
	return model.UsuarioEncrypt{
		CPF:         Encrypt(usuario.CPF, usuario.CPF),
		Nome:        Encrypt(usuario.Nome, usuario.CPF),
		Endereco:    usuario.Endereco,
		ListDividas: usuario.ListDividas,
	}
}

func Encrypt(valor string, cpf string) []byte {
	CRYPT_KEY := os.Getenv("CRYPT_KEY")
	key := []byte(cpf + CRYPT_KEY)
	plaintext := []byte(valor)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, aesgcm.NonceSize())

	ciphertext := aesgcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext
}

func DecryptUsuario(usuario model.UsuarioEncrypt, cpf string) model.UsuarioDecrypt {
	return model.UsuarioDecrypt{
		CPF:         Decrypt(usuario.CPF, cpf),
		Nome:        Decrypt(usuario.Nome, cpf),
		Endereco:    usuario.Endereco,
		ListDividas: usuario.ListDividas,
	}
}

func Decrypt(ciphertext []byte, cpf string) string {
	CRYPT_KEY := os.Getenv("CRYPT_KEY")
	key := []byte(cpf + CRYPT_KEY)
	block, _ := aes.NewCipher(key)
	// if err != nil {
	// 	panic(err.Error())
	// }
	aesgcm, _ := cipher.NewGCM(block)
	// if err != nil {
	// 	panic(err.Error())
	// }
	plaintext, _ := aesgcm.Open(nil, ciphertext[:aesgcm.NonceSize()], ciphertext[aesgcm.NonceSize():], nil)
	// if err != nil {
	// 	panic(err.Error())
	// }

	return string(plaintext)
}
