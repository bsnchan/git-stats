package main

import (
	"flag"
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
	flag.StringVar(&token, "token", "", "github token")
	flag.StringVar(&org, "org", "", "github org")
	flag.StringVar(&repo, "repo", "", "github repo")

	flag.Parse()
}
