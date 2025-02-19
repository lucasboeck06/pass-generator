package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

// Func para criptografar a senha
func Encrypt(key []byte) {

	text := CreatePassword()

	// Cria um bloco de cifra utilizando a key
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Erro ao criar cifra: %v", key)
	}

	// Cria um slice para armazenar a senha criptografada + IV
	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize] // Define o IV nos primeiros bytes do slice

	// Preenche o IV com valores aleatórios
	_, err = rand.Read(iv)
	if err != nil {
		log.Fatalf("Erro ao preencher IV: %v", err)
	}

	// Cria um stream de cifragem no modo CFB
	stream := cipher.NewCFBEncrypter(block, iv)
	// Criptografa e armazena o result depois de IV, dentro de ciphertext
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(text))

	// Transforma os bytes em uma string, legível
	// URLEncoding garante que o texto pode ser usado em URLs
	fmt.Println(base64.URLEncoding.EncodeToString(ciphertext))

}

// Func para descriptografar a senha
func Decrypt(encryptedText string, key []byte) string {

	// Decodifica o texto de base64 para bytes
	ciphertext, err := base64.URLEncoding.DecodeString(encryptedText)
	if err != nil {
		log.Fatalf("Erro ao decodificar de base64 -> bytes: %v", err)
	}

	// Cria um bloco de cifra com a chave
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Erro ao criar bloco de cifra: %v", err)
	}

	// Garante que os tamanhos batem
	if len(ciphertext) < aes.BlockSize {
		log.Fatalf("Erro, texto muito pequeno!")
	}
	iv := ciphertext[:aes.BlockSize]        // Recupera o IV (primeiros bytes de ciphetext)
	ciphertext = ciphertext[aes.BlockSize:] // define o ciphertext (tudo após o IV)

	// Cria um stream de descifragem com o CFB
	stream := cipher.NewCFBDecrypter(block, iv)
	// Transforma o ciphertext novamente para o texto original
	stream.XORKeyStream(ciphertext, ciphertext)

	// Retorna a senha descriptografada, como uma string
	return string(ciphertext)

}
