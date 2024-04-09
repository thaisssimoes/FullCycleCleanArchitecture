package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	conn *sqlx.DB
}

//TODO: ver referencia

func (db DB) OpenConnection(host, port, user, pswd, dbname string) error {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, pswd, dbname)

	conn, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return err
	}
	db.conn = conn

	err = db.conn.Ping()
	if err != nil {
		panic(err)
	}

	return nil
}

func (db DB) CloseConnection() error {
	err := db.conn.Close()
	if err != nil {
		return err
	}

	return nil
}
