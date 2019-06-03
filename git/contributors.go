package git

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Contributor struct {
	Id            int
	Login         string
	Contributions int
	Email         string
	Company       string
	Orgs          []string
}

type User struct {
	Id      int
	Login   string
	Email   string
	Company string
}

func (g *Client) GetContributors() ([]Contributor, error) {
	contributorsEndpoint := fmt.Sprintf("/repos/%s/%s/contributors", g.Org, g.Repo)
	resp, err := g.MakeRequest(contributorsEndpoint)
	if err != nil {
		return nil, err
	}

	var contributors []Contributor
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(body, &contributors)
	return contributors, nil
}

func (g *Client) GetContributorsDetailed() ([]Contributor, error) {
	contributors, err := g.GetContributors()
	if err != nil {
		return nil, err
	}

	for i, c := range contributors {
		usersEndpoint := fmt.Sprintf("/users/%s", c.Login)
		resp, err := g.MakeRequest(usersEndpoint)
		if err != nil {
			return nil, err
		}

		var user User
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err.Error())
		}
		json.Unmarshal(body, &user)
		contributors[i].Email = user.Email
		contributors[i].Company = user.Company

		userOrgsEndpoint := fmt.Sprintf("/users/%s/orgs", c.Login)
		resp, err = g.MakeRequest(userOrgsEndpoint)
		if err != nil {
			return nil, err
		}

		var orgs []struct {
			Login string
		}
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err.Error())
		}
		json.Unmarshal(body, &orgs)

		var orgsArray []string
		for _, o := range orgs {
			orgsArray = append(orgsArray, o.Login)
		}
		contributors[i].Orgs = orgsArray
	}
	return contributors, nil
}
