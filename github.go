package main

import (
	"container/ring"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"sync"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

// <URL>; rel="foobar", …
var reLink = regexp.MustCompile("<([^\\s>]+)>[\\s]*;[\\s]*rel=\\\"([^\\\"]+)\\\"")

func ParseLinkHeader(header string) (targets map[string]string) {
	targets = make(map[string]string)
	for _, match := range reLink.FindAllStringSubmatch(header, -1) {
		targets[match[2]] = match[1]
	}
	return
}

func GetJSONData(url string, data interface{}) (links map[string]string, err error) {
	resp, err := httpClient.Get(url)
	if err != nil {	return }

	links = ParseLinkHeader(resp.Header.Get("Link"))

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {	return }

	err = json.Unmarshal(body, &data)
	return
}

type GithubRepo struct {
	Languages    map[string]int
	FullName     string `json:"full_name"`
	HTMLURL      string `json:"html_url"`
	LanguagesURL string `json:"languages_url"`
}

type GithubRepoSearch struct {
	TotalCount        int          `json:"total_count"`
	IncompleteResults bool         `json:"incomplete_results"`
	Items             []GithubRepo `json:"items"`
}

func SearchGithub(query string, limit int) (results []GithubRepo, err error) {
	results = make([]GithubRepo, 0, limit)

	u, err := url.Parse("https://api.github.com/search/repositories?sort=updated&order=desc")
	if err != nil {
		return
	}

	q := u.Query()
	q.Set("q", query)
	u.RawQuery = q.Encode()

	var searchURL = u.String()
	var apiData GithubRepoSearch
	var links map[string]string
	var ok = true

	for ok && len(results) < limit {
		links, err = GetJSONData(searchURL, &apiData)
		searchURL, ok = links["next"]

		if err != nil {
			return
		}

		for _, item := range apiData.Items {
			n := len(results)
			if n == limit {
				break
			}
			results = results[:n+1]
			results[n] = item
		}
	}

	return
}

func (repo *GithubRepo) GetLanguagesData() (err error) {
	_, err = GetJSONData(repo.LanguagesURL, &repo.Languages)
	return
}

var WorkerPoolSize int
var WorkerPool *ring.Ring
var WorkerWg sync.WaitGroup

func Worker(repos chan *GithubRepo) {
	for repo := range repos {
		repo.GetLanguagesData()
	}
	WorkerWg.Done()
}

func MakeWorker() (channel chan *GithubRepo) {
	channel = make(chan *GithubRepo)
	go Worker(channel)
	return
}

func CreateWorkers() {
	r := ring.New(WorkerPoolSize)
	for i := 0; i < WorkerPoolSize; i++ {
		r.Value = MakeWorker()
		r = r.Next()
	}

	WorkerPool = r
	WorkerWg.Add(WorkerPoolSize)
}

func DispatchRepo(repo *GithubRepo) {
	worker := WorkerPool.Value.(chan *GithubRepo)
	worker <- repo
}

func CloseChannelsAndWait() {
	WorkerPool.Do(func (value interface{}) {
		channel := value.(chan *GithubRepo)
		if channel != nil {
			close(channel)
		}
	})
	WorkerWg.Wait()
}
