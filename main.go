package main

import (
	"go_gin_gorm_rest/db"
)

func main() {
	db.Init()
	db.Close()
}
