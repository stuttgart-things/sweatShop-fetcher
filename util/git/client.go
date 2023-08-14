/*
Copyright © 2023 Xiaomin Lai
*/

package git

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	memfs "github.com/go-git/go-billy/v5/memfs"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	transport "github.com/go-git/go-git/v5/plumbing/transport"
	memory "github.com/go-git/go-git/v5/storage/memory"
)

var (
	maxAttemptsCount = 1
	// maxRetryDuration time.Duration
	// retryDuration    time.Duration
	// factor           int64
)

// Client is a generic git client interface
type Client interface {
	Root() string
	Init() error
	Fetch(revision string) error
	Checkout(revision string) error
	LsFiles(path string) ([]string, error)
	TestRepository() error
}

// nativeGitClient implements Client interface with extra handling using go-git
type nativeGitClient struct {
	// URL of the repository
	repoURL string
	// Root path of repository
	root string
	// Authenticator credentials for private repositories
	creds transport.AuthMethod
	// Whether to connect insecurely to repository, e.g. don't verify certificate
	insecure bool
	// repository from go-git
	repo *git.Repository
}

// NewClient creates a new git client
func NewClient(repoURL string, root string, creds Creds, insecure bool) (Client, error) {
	return &nativeGitClient{
		repoURL:  repoURL,
		root:     root,
		creds:    creds,
		insecure: insecure,
	}, nil
}

// Init initializes a local git repository and sets the remote origin
func (m *nativeGitClient) Init() error {
	log.Infof("nativeGitClient Init was invoked for %s", m.repoURL)

	r, err := git.PlainOpen(m.root)
	if err == nil {
		m.repo = r
		log.Infoln("Repository already exists, skipping init")
		return nil
	}
	if err != git.ErrRepositoryNotExists {
		return err
	}

	log.Infof("Initializing %s to %s", m.repoURL, m.root)
	_, err = exec.Command("rm", "-rf", m.root).Output()
	if err != nil {
		return fmt.Errorf("unable to clean repo at %s: %v", m.root, err)
	}
	err = os.MkdirAll(m.root, 0755)
	if err != nil {
		return err
	}
	repo, err := git.PlainClone(m.root, false, &git.CloneOptions{
		URL:  m.repoURL,
		Auth: m.creds,
	})
	if err != nil {
		return err
	}

	m.repo = repo
	return nil
}

func (m *nativeGitClient) Root() string {
	return m.root
}

func (m *nativeGitClient) Fetch(revision string) error {
	log.Infof("Fetching %s (revision '%s') to %s", m.repoURL, revision, m.root)

	log.Debugf("Debug fetching input %+v", m)

	var err error
	if revision != "" {
		err = m.repo.Fetch(&git.FetchOptions{
			RemoteName:      git.DefaultRemoteName,
			RefSpecs:        []config.RefSpec{config.RefSpec(fmt.Sprintf("+refs/heads/%s:refs/remotes/origin/%s", revision, revision))},
			Auth:            m.creds,
			Tags:            git.AllTags,
			Force:           true,
			InsecureSkipTLS: m.insecure,
		})

		if err == git.NoErrAlreadyUpToDate {
			log.Infof("Already up-to-date")
			return nil
		}
	} else {
		err = m.repo.Fetch(&git.FetchOptions{
			RemoteName:      git.DefaultRemoteName,
			Auth:            m.creds,
			Tags:            git.AllTags,
			Force:           true,
			InsecureSkipTLS: m.insecure,
		})

		if err == git.NoErrAlreadyUpToDate {
			log.Infof("Already up-to-date")
			return nil
		}
	}
	return err
}

// Checkout checkout specified revision
func (m *nativeGitClient) Checkout(revision string) error {
	log.Infof("Checking out %s (revision '%s') to %s", m.repoURL, revision, m.root)

	if revision == "" || revision == "HEAD" {
		ref, err := m.repo.Head()
		if err != nil {
			return fmt.Errorf("Unable to get HEAD: %v", err)
		}
		revision = ref.Name().String()
		log.Infof("Using HEAD revision %s", revision)
	}

	revisionHash, _ := m.repo.ResolveRevision(plumbing.Revision(revision))
	log.Debugf("revision ('%s') hash: %s", revision, revisionHash.String())

	worktree, err := m.repo.Worktree()
	if err != nil {
		return err
	}
	if worktree == nil {
		return fmt.Errorf("worktree is nil")
	}

	err = worktree.Checkout(&git.CheckoutOptions{
		Force: true,
		Hash:  *revisionHash,
	})
	if err != nil {
		return err
	}

	// git clean -f -d
	err = worktree.Clean(&git.CleanOptions{})
	if err != nil {
		return err
	}

	return nil
}

// LsFiles lists the local working tree, including only files that are under source control
func (m *nativeGitClient) LsFiles(path string) ([]string, error) {
	log.Infof("Listing files with pattern %s", path)

	out := make([]string, 0)

	r := m.repo
	log.Debugf("output repo: %#v", r)

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	log.Debugf("output ref: %#v", ref)
	if err != nil {
		return out, err
	}

	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	log.Tracef("output commit: %#v", commit)
	if err != nil {
		return out, err
	}

	// ... retrieve the tree from the commit
	tree, err := commit.Tree()
	log.Tracef("output tree: %#v", tree)
	if err != nil {
		return out, err
	}

	// ... get the files iterator
	err = tree.Files().ForEach(func(f *object.File) error {
		if yes, err := filepath.Match(path, f.Name); yes && err == nil {
			log.Debugf("file: %s, hash: %s", f.Name, f.Hash)
			out = append(out, f.Name)
		}
		return nil
	})
	if err != nil {
		return out, err
	}

	return out, nil
}

// TestRepository tests the repository connection and authentication
func (c *nativeGitClient) TestRepository() error {
	// Init memory storage and fs
	storer := memory.NewStorage()
	fs := memfs.New()

	// Clone repo into memfs
	_, err := git.Clone(storer, fs, &git.CloneOptions{
		URL:  c.repoURL,
		Auth: c.creds,
	})
	if err != nil {
		return fmt.Errorf("Could not git clone repository %s: %w\n", c.repoURL, err)
	}
	log.Infoln("Repository cloned")

	return nil
}
