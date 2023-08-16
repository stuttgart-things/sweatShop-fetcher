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
	testRepository(c)
}

func testRepository(c apiclient.FetcherServiceClient) {
	log.Println("Example: test repository connection")

	// Create a new repository
	repo := &apiclient.Repository{
		Repo: "https://github.com/argoproj/argo-cd",
		// Repo:     "https://codehub.sva.de/Lab/stuttgart-things/dev/sthingsK8s.git",
		// Username: "<CODEHUB-USERNAME>",
		// Password: "<CODEHUB-ACCESS-TOKEN>",
	}

	res, err := c.TestRepository(context.Background(), &apiclient.TestRepositoryRequest{Repo: repo})
	if err != nil {
		log.Fatalf("could not test repository: %v", err)
	}

	log.Printf("TestRepository output: %v\n", res)
}
