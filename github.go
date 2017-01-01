package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func GetJSONData(url string, data interface{}) (err error) {
	resp, err := httpClient.Get(url)
	if err != nil {	return }

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {	return }

	err = json.Unmarshal(body, &data)
	return
}

type GithubRepo struct {
	Languages map[string]int
	FullName string `json:"full_name"`
	HTMLURL string `json:"html_url"`
	LanguagesURL string `json:"languages_url"`
}

type GithubRepoSearch struct {
	TotalCount int `json:"total_count"`
	IncompleteResults bool `json:"incomplete_results"`
	Items []GithubRepo `json:"items"`
}


func SearchGithub(query string) (results []GithubRepo, err error) {
	results= make([]GithubRepo, 0, 100)
	
	u, err := url.Parse("https://api.github.com/search/repositories?sort=updated&order=desc")
	if err != nil { return }

	q := u.Query()
	q.Set("q", query)
	u.RawQuery = q.Encode()	

	var apiData GithubRepoSearch
	err = GetJSONData(u.String(), &apiData)
	if err != nil { return }

	for _, item := range apiData.Items {
		n := len(results)
		results = results[:n+1]
		results[n] = item
	}

	return 
}
