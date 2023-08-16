package main

import (
	"context"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"

	"github.com/stuttgart-things/sweatShop-fetcher/fetcher/apiclient"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := apiclient.NewFetcherServiceClient(conn)

	// Contact the server and print out its response.
	listFiles(c)
}

func listFiles(c apiclient.FetcherServiceClient) {
	log.Println("Example: list files")

	// Create a new repository
	repo := &apiclient.Repository{
		Repo: "https://github.com/argoproj/argo-cd",
	}

	// Create a new list files request
	in := &apiclient.ListFilesRequest{
		Repo:     repo,
		Revision: "release-1.0",
		Path:     "go.*",
	}

	res, err := c.ListFiles(context.Background(), in)
	if err != nil {
		log.Fatalf("could not list files: %v", err)
	}

	log.Printf("ListFiles output: %v\n", res)
}
