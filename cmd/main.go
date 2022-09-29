package main

import (
	"context"
	"github-actions/cmd/app"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func main() {

	hostAddr := os.Getenv("HOST_ADDR")
	if hostAddr == "" {
		hostAddr = "127.0.0.1"
	}

	hostPortStr := os.Getenv("HOST_PORT")
	if hostPortStr == "" {
		hostPortStr = "8080"
	}

	hostPort, err := strconv.Atoi(hostPortStr)
	if err != nil {
		panic(err)
	}

	app := app.App{HostAddr: hostAddr, HostPort: hostPort}
	go app.Start()

	// Setting up signal capturing
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Waiting for SIGINT (kill -2)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	app.Stop(ctx)

}
