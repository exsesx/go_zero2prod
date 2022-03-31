package main

import (
	"go_zero2prod/config"
	"go_zero2prod/server"
)

func main() {
	config.Init()
	server.Init()
}
