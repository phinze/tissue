package tissue

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func HandleIssue(owner, repo, issueId string) {
	issueNum, err := strconv.Atoi(issueId)
	if err != nil {
		log.Fatalf("err: %s", err)
	}

	issue, _, err := githubClient().Issues.Get(owner, repo, issueNum)
	if err != nil {
		log.Fatalf("err: %s", err)
	}
	if issue == nil {
		log.Fatalf("issue came back nil!")
	}

	issueSlug := slugify(strings.Join([]string{issueId, *issue.Title}, "-"))
	log.Printf("issue slug: %s", issueSlug)

	issueDir := filepath.Join(issuesRoot(), "github.com", owner, repo, issueSlug)
	log.Printf("issue dir: %s", issueDir)

	if err := os.MkdirAll(issueDir, 0755); err != nil {
		log.Fatalf("err: %s", err)
	}

	issueMd, err := os.Create(filepath.Join(issueDir, "issue.markdown"))
	if err != nil {
		log.Fatalf("err: %s", err)
	}
	defer issueMd.Close()
	issueMd.WriteString(fmt.Sprintf("# %s (#%d)\n\n", *issue.Title, *issue.Number))
	issueMd.WriteString(*issue.Body + "\n\n")

	issueUrl, err := os.Create(filepath.Join(issueDir, "issue.url"))
	if err != nil {
		log.Fatalf("err: %s", err)
	}
	defer issueUrl.Close()
	issueUrl.WriteString(*issue.URL + "\n")

	if err := os.Chdir(issueDir); err != nil {
		log.Fatalf("err: %s", err)
	}

	fmt.Println(issueDir)
}
