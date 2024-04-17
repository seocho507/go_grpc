package main

import (
	"flag"
	"fmt"
	"go_grpc/cmd"
	"go_grpc/config"
	"go_grpc/grpc/server"
	"log"
	"time"
)

var configFlag = flag.String("config", "./config.toml", "path to config file")

func main() {
	flag.Parse()

	cnf, err := config.NewConfig(*configFlag)
	if err != nil {
		fmt.Println("failed to load config:", err)
		return
	}
	if _, err := server.NewGRPCServer(cnf); err != nil {
		log.Fatal(err)
	}
	time.Sleep(1e9)
	_, err = cmd.NewApp(cnf)
	if err != nil {
		log.Fatal(err)
	}
}
