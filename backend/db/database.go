package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool // Pool de conexão global

func InitDatabase() {
	// Carrega as variáveis da .env
	err := godotenv.Load("../../deploy/.env")
	if err != nil {
		log.Fatalf(".env não pode ser acessada: %v", err)
	}

	// Configuração do banco
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}

	DB = pool
	log.Println("Conectado ao banco de dados com sucesso!")
}

func CloseDatabase() {
	if DB != nil {
		DB.Close()
		log.Println("Conexão com banco encerrada.")
	}
}
