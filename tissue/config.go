package tissue

import (
	"log"
	"os"
	"path/filepath"

	"github.com/google/go-github/github"
	"github.com/mitchellh/go-homedir"
)

func issuesRoot() string {
	issuesRoot := os.Getenv("TISSUE_ISSUES_ROOT")
	if issuesRoot == "" {
		issuesRoot = "~/issues"
	}
	issuesRoot, err := homedir.Expand(issuesRoot)
	if err != nil {
		log.Fatalf("err: %s", err)
	}
	return issuesRoot
}

func srcRoot(repo *github.Repository) string {
	root := "~/src"
	if repo.Language != nil && *repo.Language == "Go" {
		root = "~/go/src"
	}
	root, err := homedir.Expand(root)
	if err != nil {
		log.Fatalf("err: %s", err)
	}
	return root
}

func repoLocalPath(repo *github.Repository) string {
	return filepath.Join(srcRoot(repo), "github.com", *repo.Owner.Login, *repo.Name)
}
