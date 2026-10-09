package mysql

import (
	"database/sql"
)

type MysqlDB struct {
	db *sql.DB
}

func New() (*MysqlDB, error) {
	dsn := "user:5678docker@tcp(127.0.0.1:3308)/quizofkingzdb"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return &MysqlDB{db: db}, nil
}
