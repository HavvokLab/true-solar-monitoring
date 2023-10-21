package main

import (
	"flag"
	"log"

	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/domain"
	"github.com/HavvokLab/true-solar-monitoring/infra"
	"github.com/HavvokLab/true-solar-monitoring/logger"
	"github.com/HavvokLab/true-solar-monitoring/repo"
	"github.com/HavvokLab/true-solar-monitoring/service"
)

var username *string = flag.String("usr", "", "username of new user")
var password *string = flag.String("pwd", "", "password of new user")

func init() {
	config.InitConfig()
}

func init() {
	infra.InitGormDB()
}

func main() {
	flag.Parse()
	repo := repo.NewUserRepo(infra.GormDB)
	serv := service.NewAuthService(repo, logger.NewLoggerMock())
	if err := serv.Register(&domain.RegisterRequest{
		Username: *username,
		Password: *password,
	}); err != nil {
		log.Fatal(err)
	}
	log.Default().Println("User registered successfully")
}
