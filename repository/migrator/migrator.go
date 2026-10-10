package migrator

import (
	"database/sql"
	"fmt"
	"quizofkingz/repository/mysql"

	migrate "github.com/rubenv/sql-migrate"
)

type Migrator struct {
	dialect    string
	dbConfig   mysql.Config
	migrations *migrate.FileMigrationSource
}

func New(dbConfig mysql.Config) Migrator {
	migrations := &migrate.FileMigrationSource{
		Dir: "./repository/mysql/migrations",
	}
	return Migrator{dbConfig: dbConfig, migrations: migrations, dialect: "mysql"}
}

func (m Migrator) Up() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", m.dbConfig.Username, m.dbConfig.Password, m.dbConfig.Host, m.dbConfig.Port, m.dbConfig.DBName)
	db, err := sql.Open(m.dialect, dsn)
	if err != nil {
		panic(fmt.Errorf("can't open db connection: %v", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("migration failed: %v", err))
	}
	fmt.Printf("%d migrations applied\n", n)
}

func (m Migrator) Down() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", m.dbConfig.Username, m.dbConfig.Password, m.dbConfig.Host, m.dbConfig.Port, m.dbConfig.DBName)
	db, err := sql.Open(m.dialect, dsn)
	if err != nil {
		panic(fmt.Errorf("can't open db connection: %v", err))
	}

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Down)
	if err != nil {
		panic(fmt.Errorf("Rollback migrations failed: %v", err))
	}
	fmt.Printf("Rollback %d migrations successfully!\n", n)
}

func (m Migrator) Status() {
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", m.dbConfig.Username, m.dbConfig.Password, m.dbConfig.Host, m.dbConfig.Port, m.dbConfig.DBName)
	//db, err := sql.Open(m.dialect, dsn)
	//if err != nil {
	//	panic(fmt.Errorf("can't open db connection: %v", err))
	//}
	//
	//n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Up)
	//if err != nil {
	//	panic(fmt.Errorf("migration failed: %v", err))
	//}
	//fmt.Printf("%d migrations applied\n", n)
}
