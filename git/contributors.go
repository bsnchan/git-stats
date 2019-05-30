package git

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Contributor struct {
	Id            int
	Login         string
	Contributions int
}

func (g *Client) GetContributors() ([]Contributor, error) {
	contributorsEndpoint := fmt.Sprintf("https://api.github.com/repos/%s/%s/contributors", g.Org, g.Repo)
	req, err := http.NewRequest("GET", contributorsEndpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("token %s", g.Token))
	resp, err := g.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		panic("Did not get 200 status code")
	}

	var contributors []Contributor
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(body, &contributors)
	return contributors, nil
}
