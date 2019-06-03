package git

import (
	"fmt"
	"net/http"
)

type Client struct {
	HttpClient *http.Client
	Org        string
	Repo       string
	Token      string
}

const GithubApi = "https://api.github.com"

func NewClient(org, repo, token string) Client {
	client := Client{}
	client.HttpClient = &http.Client{}
	client.Org = org
	client.Repo = repo
	client.Token = token
	return client
}

func (g *Client) MakeRequest(endpoint string) (*http.Response, error) {
	requestEndpoint := fmt.Sprintf("%s%s", GithubApi, endpoint)
	req, err := http.NewRequest("GET", requestEndpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("token %s", g.Token))
	resp, err := g.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
