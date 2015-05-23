package tissue

import (
	"net/http"
	"os"

	"github.com/google/go-github/github"
	"github.com/phinze/tissue/basicauth"
)

func githubClient() *github.Client {
	token := os.Getenv("TISSUE_GITHUB_TOKEN")
	user := os.Getenv("TISSUE_GITHUB_USER")

	var transport http.RoundTripper
	var client *http.Client
	if token != "" && user != "" {
		transport = &basicauth.Transport{
			Username: user,
			Password: token,
		}
		client = &http.Client{
			Transport: transport,
		}
	}
	return github.NewClient(client)
}
