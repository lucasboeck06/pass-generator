package services

import (
	"crypto/rand"
	"log"
)

func CreatePassword() string {

	// Definindo os caractéres da senha
	charset := ",.:;[{]}^~`abcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+"

	// Criando slice com bytes aleatórios
	password := make([]byte, 16)  // Cria um slice com o tamanho 16
	_, err := rand.Read(password) // Preenche o slice com um valor random em byte
	if err != nil {
		log.Fatalf("Erro ao gerar bytes aleatórios: %v", err)
	}

	// Trocando os bytes por valores existentes em charset
	for i := range password {
		password[i] = charset[int(password[i])%len(charset)]
	}

	// Conversão para string e retorno da senha
	return string(password)

}
