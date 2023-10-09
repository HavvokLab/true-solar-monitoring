package main

import (
	"fmt"

	"github.com/HavvokLab/true-solar-monitoring/config"
	"github.com/HavvokLab/true-solar-monitoring/constant"
	"github.com/HavvokLab/true-solar-monitoring/infra"
	"github.com/HavvokLab/true-solar-monitoring/repo"
)

func init() {
	config.InitConfig()
}

func init() {
}

func main() {
	db, err := infra.NewGormDB()
	if err != nil {
		panic(err)
	}

	repo := repo.NewHuaweiCredentialRepo(db)
	data, err := repo.GetCredentialsByOwner(constant.TRUE_OWNER)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
