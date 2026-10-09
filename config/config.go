package config

import (
	"quizofkingz/repository/mysql"
	"quizofkingz/service/authservice"
)

type HTTPServer struct {
	Port int
}

type Config struct {
	HTTPServer HTTPServer
	Auth       authservice.Config
	Mysql      mysql.Config
}
