package git

import "net/http"

type Client struct {
	HttpClient *http.Client
	Org        string
	Repo       string
	Token      string
}

func NewClient(org, repo, token string) Client {
	client := Client{}
	client.HttpClient = &http.Client{}
	client.Org = org
	client.Repo = repo
	client.Token = token
	return client
}
