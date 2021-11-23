package main

import (
	"go_gin_gorm_rest/db"
	"go_gin_gorm_rest/server"
)

func main() {
	db.Init()
	server.Init()
	db.Close()
}
