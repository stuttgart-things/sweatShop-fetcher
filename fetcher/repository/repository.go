package repository

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"

	gogit "github.com/go-git/go-git/v5"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/stuttgart-things/sweatShop-fetcher/fetcher/apiclient"
	"github.com/stuttgart-things/sweatShop-fetcher/util/git"
	"github.com/stuttgart-things/sweatShop-fetcher/util/io"
)

type Service struct {
	rootDir      string
	newGitClient func(url string, root string, creds git.Creds, insecure bool) (git.Client, error)
	gitRepoPaths *io.TempPaths
	// TODO: add more fields. E.g.: repoLock, cache, gitRepoInitializer, etc.
}

// NewService creates a new repository service
func NewService(rootDir string) *Service {
	log.Infof("Create NewService rootDir: %s", rootDir)

	return &Service{
		rootDir:      rootDir,
		newGitClient: git.NewClient,
		gitRepoPaths: io.NewTempPaths(rootDir),
	}
}

func (s *Service) Init() error {
	log.Infof("Init new repository service")

	_, err := os.Stat(s.rootDir)
	if os.IsNotExist(err) {
		return os.MkdirAll(s.rootDir, 0300)
	}
	if err == nil {
		// give itself read permissions to list previously written directories
		err = os.Chmod(s.rootDir, 0700)
	}

	var files []fs.DirEntry
	if err == nil {
		files, err = os.ReadDir(s.rootDir)
	}
	if err != nil {
		log.Warnf("Failed to restore cloned repositories paths: %v", err)
		return nil
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		fullPath := filepath.Join(s.rootDir, file.Name())

		if repo, err := gogit.PlainOpen(fullPath); err == nil {
			if remotes, err := repo.Remotes(); err == nil && len(remotes) > 0 && len(remotes[0].Config().URLs) > 0 {
				s.gitRepoPaths.Add(git.NormalizeGitURL(remotes[0].Config().URLs[0]), fullPath)
			}
		}
	}

	log.Debugf("Restored cloned repositories paths: %+v", s.gitRepoPaths)

	// remove read permissions since no-one should be able to list the directories
	return os.Chmod(s.rootDir, 0300)
}

// ListFiles lists the contents of a GitHub repo
func (s *Service) ListFiles(ctx context.Context, q *apiclient.ListFilesRequest) (*apiclient.FileList, error) {
	log.Infof("ListFiles repo: %s, revision: %s", q.GetRepo().Repo, q.GetRevision())

	res := apiclient.FileList{}

	// Create git client
	gitClient, err := s.newClient(q.GetRepo())
	if err != nil {
		return &res, err
	}

	// TODO: add cache. Try to get files from cache. If not found, continue with the rest of the function

	// TODO: add lock. Try to lock the repo and return a closer.

	// Init repo. If repo is already initialized, this is a no-op
	err = initRepo(gitClient)
	if err != nil {
		return &res, err
	}

	// Checkout revision
	err = checkoutRevision(gitClient, q.GetRevision())
	if err != nil {
		return &res, err
	}

	// List files that match the given pattern
	files, err := gitClient.LsFiles(q.GetPath())
	res.Files = files

	return &res, nil
}

// TestRepository tests the repository connection and authentication
func (s *Service) TestRepository(ctx context.Context, in *apiclient.TestRepositoryRequest) (*apiclient.TestRepositoryResponse, error) {
	r := in.GetRepo()

	log.Infof("Test repository %s", r.Repo)

	// Create git client
	gitClient, err := s.newClient(r)
	if err != nil {
		return nil, err
	}

	// Test repository
	err = gitClient.TestRepository()
	if err != nil {
		return nil, err
	}

	apiResp := &apiclient.TestRepositoryResponse{VerifiedRepository: true}

	return apiResp, nil
}

// newClient creates a new git client
func (s *Service) newClient(r *apiclient.Repository) (git.Client, error) {
	// Create git credentials
	creds := git.NewHTTPBasicAuthCred(r.GetUsername(), r.GetPassword(), r.GetInsecure())

	// Get repo path
	repoPath, err := s.gitRepoPaths.GetPath(git.NormalizeGitURL(r.Repo))
	if err != nil {
		return nil, err
	}

	// Create git client
	client, err := s.newGitClient(r.Repo, repoPath, creds, r.GetInsecure())
	if err != nil {
		return nil, err
	}

	return client, nil
}

// initRepo initializes a git repository
func initRepo(gitClient git.Client) error {
	log.Infof("Initialize git repo %s", gitClient.Root())
	err := gitClient.Init()
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to initialize git repo: %v", err)
	}
	return nil
}

// checkoutRevision checks out the given revision
func checkoutRevision(gitClient git.Client, revision string) error {

	// Fetching with no revision first. Fetching with an explicit version can cause repo bloat. https://github.com/argoproj/argo-cd/issues/8845
	err := gitClient.Fetch("")
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to fetch default: %v", err)
	}

	err = gitClient.Checkout(revision)
	if err != nil {
		// When fetching with no revision, only refs/heads/* and refs/remotes/origin/* are fetched. If checkout fails
		// for the given revision, try explicitly fetching it.
		log.Infof("Failed to checkout revision %s: %v", revision, err)
		log.Infof("Fallback to fetching specific revision %s. ref might not have been in the default refspec fetched.", revision)

		err = gitClient.Fetch(revision)
		if err != nil {
			return status.Errorf(codes.Internal, "Failed to checkout revision %s: %v", revision, err)
		}

		err = gitClient.Checkout("FETCH_HEAD")
		if err != nil {
			return status.Errorf(codes.Internal, "Failed to checkout FETCH_HEAD: %v", err)
		}
	}

	return err
}
