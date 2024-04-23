package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	Conn *sqlx.DB
}

//TODO: ver referencia

func OpenConnection(host, user, pswd, dbname string, port int) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pswd, dbname)

	conn, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func CloseConnection(db *sqlx.DB) {
	db.Close()
}
