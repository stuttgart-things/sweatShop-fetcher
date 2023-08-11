package git

import (
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

type Creds interface {
	transport.AuthMethod
}

// NewHTTPBasicAuthCred creates a new HTTP basic authentication cred
func NewHTTPBasicAuthCred(username string, password string, insecure bool) *http.BasicAuth {
	if username == "" || password == "" {
		return nil
	}
	return &http.BasicAuth{
		Username: username,
		Password: password,
	}
}
