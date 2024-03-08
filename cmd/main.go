package main

import (
	"ChargPiles/config"
	"ChargPiles/repository/db/dao"
	"ChargPiles/routes"
)

func main() {
	loading()

	r := routes.NewRouter()
	_ = r.Run(":8080")
}

func loading() {
	config.Init()
	dao.MySQLInit()
}