package tissue

import (
	"log"
	"os"

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
