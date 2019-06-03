package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/bsnchan/git-stats/git"
)

var token string
var org string
var repo string

func main() {
	g := git.NewClient(org, repo, token)
	contributors, err := g.GetContributorsDetailed()
	if err != nil {
		panic(err.Error())
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t", "LOGIN", "CONTRIBUTIONS", "EMAIL", "COMPANY", "ORGS")
	for _, c := range contributors {
		orgs := strings.Join(c.Orgs[:], ",")
		fmt.Fprintf(w, "\n %s\t%d\t%s\t%s\t%s\t", c.Login, c.Contributions, c.Email, c.Company, orgs)
	}
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
