package tissue

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func HandleRepo(owner, repo string) {
	r, _, err := githubClient().Repositories.Get(owner, repo)
	if err != nil {
		log.Fatalf("err: %s", err)
	}

	localPath := repoLocalPath(r)

	if _, err := os.Stat(localPath); os.IsNotExist(err) {
		if err := exec.Command("git", "clone", *r.CloneURL, localPath).Run(); err != nil {
			log.Fatalf("err: %s", err)
		}
	}

	fmt.Println(localPath)
}
