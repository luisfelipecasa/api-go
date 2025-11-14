package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "localhost"
	port     = "3306"
	user     = "root"
	password = ""
	dbname   = "api_go"
)

func ConnectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		user, password, host, port, dbname,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Erro ao abrir conexão:", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("Erro ao pingar:", err)
		return nil, err
	}

	fmt.Println("Conectado ao MySQL:", dbname)

	return db, nil
}
