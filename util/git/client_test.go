package git

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/stuttgart-things/sweatShop-fetcher/fetcher/apiclient"
)

func TestInit(t *testing.T) {
	dir := t.TempDir()
	repoPath := filepath.Join(dir, "argo-cd")
	r := &apiclient.Repository{
		Repo: "https://github.com/argoproj/argo-cd",
	}

	gitClient, err := NewClient(r.Repo, repoPath, nil, r.GetInsecure())
	assert.NoError(t, err)
	assert.NotNil(t, gitClient)
	require.NoError(t, gitClient.Init())
	assert.Equal(t, gitClient.Root(), repoPath)
}

func TestFetch(t *testing.T) {
	dir := t.TempDir()
	repoPath := filepath.Join(dir, "argo-cd")
	r := &apiclient.Repository{
		Repo: "https://github.com/argoproj/argo-cd",
	}

	gitClient, err := NewClient(r.Repo, repoPath, nil, r.GetInsecure())
	assert.NoError(t, err)
	assert.NotNil(t, gitClient)
	require.NoError(t, gitClient.Init())
	require.NoError(t, gitClient.Fetch(""))
	require.NoError(t, gitClient.Fetch("release-2.6"))
}

func TestLsFiles(t *testing.T) {
	dir := t.TempDir()
	repoPath := filepath.Join(dir, "argo-cd")
	r := &apiclient.Repository{
		Repo: "https://github.com/argoproj/argo-cd",
	}

	gitClient, err := NewClient(r.Repo, repoPath, nil, r.GetInsecure())
	assert.NoError(t, err)
	assert.NotNil(t, gitClient)
	require.NoError(t, gitClient.Init())
	require.NoError(t, gitClient.Checkout(""))
	require.NoError(t, gitClient.Checkout("origin/release-2.6"))
	files, err := gitClient.LsFiles("go.*")
	assert.NoError(t, err)
	assert.NotEmpty(t, files)
	assert.Contains(t, files, "go.mod")
	assert.Contains(t, files, "go.sum")
	assert.NotContains(t, files, "Makefile")
}
