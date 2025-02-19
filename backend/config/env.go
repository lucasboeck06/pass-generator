package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Func de carregamento da .env
func LoadEnv() {
	// Carrega as variáveis de ambiente
	err := godotenv.Load("../../deploy/.env")
	if err != nil {
		log.Fatal("Erro ao carregar .env!")
	}
}

// Func para o carregamento da chave de criptografia
func GetKey() {
	// Carrega a KEY e atribui a key
	fmt.Println(os.Getenv("KEY"))
}

// Func para o carregamento das variáveis do DB
func GetDBConfig() map[string]string {
	return map[string]string{
		"host":    os.Getenv("DB_HOST"),
		"port":    os.Getenv("DB_PORT"),
		"user":    os.Getenv("DB_USER"),
		"pasword": os.Getenv("DB_PASSWORD"),
		"name":    os.Getenv("DB_NAME"),
	}
}
