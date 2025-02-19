package db

import (
	"context"
	"fmt"
	"log"
	"pass-generator/backend/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool // Pool de conexão global

func InitDatabase() {

	dbConfig := config.GetDBConfig()

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbConfig["user"], dbConfig["password"], dbConfig["host"], dbConfig["port"], dbConfig["name"])
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
