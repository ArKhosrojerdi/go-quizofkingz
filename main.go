package main

import (
	"fmt"
	"quizofkingz/config"
	"quizofkingz/delivery/httpserver"
	"quizofkingz/pkg/auth"
	"quizofkingz/repository/migrator"
	"quizofkingz/repository/mysql"
	"quizofkingz/service/authservice"
	"quizofkingz/service/userservice"
	"time"
)

const (
	JwtSignKey                     = "sign"
	AccessTokenSubject             = "at"
	RefreshTokenSubject            = "rt"
	AccessTokenExpirationDuration  = time.Hour * 24
	RefreshTokenExpirationDuration = time.Hour * 24 * 7
)

func main() {
	cnf := config.Config{
		HTTPServer: config.HTTPServer{Port: 8080},
		Auth: authservice.Config{
			SignKey:                   JwtSignKey,
			AccessSubject:             AccessTokenSubject,
			RefreshSubject:            RefreshTokenSubject,
			AccessExpirationDuration:  AccessTokenExpirationDuration,
			RefreshExpirationDuration: RefreshTokenExpirationDuration,
		},
		Mysql: mysql.Config{
			Host:     "localhost",
			Port:     3308,
			Username: "user",
			Password: "5678docker",
			DBName:   "quizofkingzdb",
		}}

	authSvc, userSvc := setupServices(cnf)

	mgtr := migrator.New(cnf.Mysql)
	mgtr.Up()
	
	server := httpserver.New(cnf, authSvc, userSvc)

	server.Serve()
}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service) {
	authSvc := authservice.New(cfg.Auth)

	MysqlRepo, err := mysql.New(cfg.Mysql)
	if err != nil {
		panic(fmt.Errorf("mysql init err: %v", err))
	}

	privateKey, err := auth.LoadPrivateKey()
	if err != nil {
		return authservice.Service{}, userservice.Service{}
	}

	userSvc := userservice.New(privateKey, MysqlRepo)

	return authSvc, userSvc
}
