package git

import (
	"os"

	gitv5 "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

// Credentials for the git client
type Credentials struct {
	UserName    string
	AccessToken string
	Repository  string
}

// Read credentials from environment variables
func GetCredentials() Credentials {
	return Credentials{
		UserName:    os.Getenv("GIT_USER_NAME"),
		AccessToken: os.Getenv("GIT_ACCESS_TOKEN"),
		Repository:  os.Getenv("LYRICS_REPOSITORY"),
	}
}

func GetRepo(dir string, credentials Credentials) error {
	_, err := gitv5.PlainClone(dir, false, &gitv5.CloneOptions{
		Auth: &http.BasicAuth{
			Username: credentials.UserName,
			Password: credentials.AccessToken,
		},
		URL:      credentials.Repository,
		Progress: os.Stdout,
	})
	if err != nil {
		return err
	}
	return nil
}
