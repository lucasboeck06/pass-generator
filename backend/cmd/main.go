package main

import (
	"pass-generator/backend/config"
)

func main() {

	// db.InitDatabase()        // Puxa a func de db/database.go e inicia o banco de dadoe
	// defer db.CloseDatabase() // Puxa a func e fecha o database (comando final)

	// stop := make(chan os.Signal, 1)
	// signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// log.Println("Servidor rodando... Preissone Ctrl+C para encerrar.")

	// <-stop
	// log.Println("Encerrando aplicação...")

	config.LoadEnv()
	config.GetKey()

}
