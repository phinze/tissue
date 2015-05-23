package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/phinze/tissue/tissue"
)

func main() {
	var shell bool
	flag.BoolVar(&shell, "sh", false, "output tj shell function")
	flag.Usage = usage
	flag.Parse()

	if shell {
		tissue.Shell()
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) < 1 {
		usage()
	}

	if os.Getenv("TISSUE_DEBUG") != "" {
		log.SetOutput(os.Stderr)
	} else {
		log.SetOutput(ioutil.Discard)
	}

	tissue.HandleURL(args[0])
}

func usage() {
	fmt.Fprintf(os.Stderr, "NAME\n\n  tissue: a github issues workflow\n\n")
	fmt.Fprintf(os.Stderr, "USAGE\n\n  tissue [OPTIONS] <GITHUB_URL>\n\n")
	fmt.Fprintf(os.Stderr, "OPTIONS\n\n")
	flag.PrintDefaults()
	os.Exit(2)
}
