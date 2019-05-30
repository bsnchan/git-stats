package main

import (
	"fmt"
	"os"

	"github.com/bsnchan/git-stats/git"
)

var token string
var org string
var repo string

func main() {
	g := git.NewClient(org, repo, token)
	fmt.Println(g)
	contributors, err := g.GetContributors()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(contributors)
}

func init() {
	if token = os.Getenv("TOKEN"); token == "" {
		panic("TOKEN must be specified")
	}
	if org = os.Getenv("ORG"); org == "" {
		panic("ORG must be specified")
	}
	if repo = os.Getenv("REPO"); repo == "" {
		panic("REPO must be specified")
	}
}
