package mysql

import (
	"database/sql"
	"fmt"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

type MysqlDB struct {
	config Config
	db     *sql.DB
}

func New(config Config) (*MysqlDB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Username, config.Password, config.Host, config.Port, config.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Errorf("cant't open db connection: %v", err))
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return &MysqlDB{config: config, db: db}, nil
}
