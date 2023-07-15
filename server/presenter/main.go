package main

import (
	"data"
	"presenter/server"
)

func main() {
	data.InitializeCrypto()
	data.AutoMigration()
	server.Server()
}
