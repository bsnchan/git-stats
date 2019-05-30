package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Contributor struct {
	Id            int
	Login         string
	Contributions int
}

var token string
var org string
var repo string

func main() {
	client := &http.Client{}
	contributorsEndpoint := fmt.Sprintf("https://api.github.com/repos/%s/%s/contributors", org, repo)
	req, err := http.NewRequest("GET", contributorsEndpoint, nil)
	if err != nil {
		panic(err.Error())
	}
	req.Header.Add("Authorization", fmt.Sprintf("token %s", token))
	resp, err := client.Do(req)
	if err != nil {
		panic(err.Error())
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
