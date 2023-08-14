/*
Copyright © 2023 Xiaomin Lai
*/

package repository

import (
	"context"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/stuttgart-things/sweatShop-fetcher/fetcher/apiclient"
	"github.com/stuttgart-things/sweatShop-fetcher/util/git"
)

func TestInit(t *testing.T) {
	dir := t.TempDir()

	// service.Init sets permission to 0300. Restore permissions when the test
	// finishes so dir can be removed properly.
	t.Cleanup(func() {
		require.NoError(t, os.Chmod(dir, 0777))
	})

	repoPath := path.Join(dir, "repo1")
	require.NoError(t, initGitRepo(repoPath, "https://github.com/argo-cd/test-repo1"))

	service := NewService(".")
	service.rootDir = dir

	require.NoError(t, service.Init())

	repo1Path, err := service.gitRepoPaths.GetPath(git.NormalizeGitURL("https://github.com/argo-cd/test-repo1"))
	assert.NoError(t, err)
	assert.Equal(t, repoPath, repo1Path)

	_, err = os.ReadDir(dir)
	require.Error(t, err)
	require.NoError(t, initGitRepo(path.Join(dir, "repo2"), "https://github.com/argo-cd/test-repo2"))
}

func Test_TestRepository(t *testing.T) {
	dir := t.TempDir()

	service := NewService(".")
	service.rootDir = dir

	in := &apiclient.TestRepositoryRequest{
		Repo: &apiclient.Repository{
			Repo: "https://github.com/argoproj/argo-cd",
		},
	}
	resp, err := service.TestRepository(context.Background(), in)

	require.NoError(t, err)
	assert.Equal(t, &apiclient.TestRepositoryResponse{
		VerifiedRepository: true,
	}, resp)
}

func initGitRepo(repoPath string, remote string) error {
	if err := os.Mkdir(repoPath, 0755); err != nil {
		return err
	}

	cmd := exec.Command("git", "init", repoPath)
	cmd.Dir = repoPath
	if err := cmd.Run(); err != nil {
		return err
	}
	cmd = exec.Command("git", "remote", "add", "origin", remote)
	cmd.Dir = repoPath
	return cmd.Run()
}
