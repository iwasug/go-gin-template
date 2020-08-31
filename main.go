package main

import (
	"flag"
	"fmt"
	"os"

	"go-gin-template/config"
	"go-gin-template/db"
	"go-gin-template/server"
)

func main() {
	env := config.GetEnv()
	environment := flag.String("e", env, "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	server.Init()
}
