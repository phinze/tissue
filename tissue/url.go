package tissue

import (
	"log"
	"net/url"
	"strings"
)

func HandleURL(argUrl string) {
	url, err := url.Parse(argUrl)
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.SplitN(url.Path[1:], "/", 4)

	switch parts[2] {
	case "issues":
		HandleIssue(parts[0], parts[1], parts[3])
	default:
		log.Fatalf("Cannot handle object type: %q", parts[2])
	}
}
