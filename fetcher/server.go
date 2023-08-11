package fetcher

import (
	"context"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/stuttgart-things/sweatShop-fetcher/fetcher/apiclient"
)

// RepoServer is the main struct for the repo server
type RepoServer struct {
	repoService *repository.Service
	opts        []grpc.ServerOption
	// TODO: add more services. E.g.: cache, etc.
}

// NewRepoServer creates a new RepoServer
func NewRepoServer(ctx context.Context) (*RepoServer, error) {
	log.Infof("Create new repo server")

	repoService := repository.NewService(filepath.Join(os.TempDir(), "_yacht-repo"))
	if err := repoService.Init(); err != nil {
		return nil, err
	}
	log.Infof("Created and initialized new repo service")

	return &RepoServer{
		repoService: repoService,
	}, nil
}

// CreateGRPC creates new configured grpc server
func (a *RepoServer) CreateGRPC() *grpc.Server {
	server := grpc.NewServer(a.opts...)

	apiclient.RegisterRepoServerServiceServer(server, a.repoService)

	// Register reflection service on gRPC server.
	reflection.Register(server)

	return server
}
