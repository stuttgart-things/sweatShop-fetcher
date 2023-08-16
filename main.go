/*
Copyright © 2023 Xiaomin Lai
*/

package main

import (
	"context"
	"github/sweatShop-fetcher/internal"

	log "github.com/sirupsen/logrus"

	"net"

	"github.com/stuttgart-things/sweatShop-fetcher/fetcher"
)

var addr string = "0.0.0.0:50051"

func main() {

	// Set the logger to debug level
	// log.SetLevel(log.DebugLevel)

	internal.PrintBanner()

	server, err := fetcher.NewRepoServer(context.Background())
	if err != nil {
		panic(err)
	}
	grpc := server.CreateGRPC()

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	if err := grpc.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
