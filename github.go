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
	ID int `json:"id"`
	Name string `json:"name"`
	FullName string `json:"full_name"`
	Owner struct {
		Login string `json:"login"`
		ID int `json:"id"`
		AvatarURL string `json:"avatar_url"`
		GravatarID string `json:"gravatar_id"`
		URL string `json:"url"`
		HTMLURL string `json:"html_url"`
		FollowersURL string `json:"followers_url"`
		FollowingURL string `json:"following_url"`
		GistsURL string `json:"gists_url"`
		StarredURL string `json:"starred_url"`
		SubscriptionsURL string `json:"subscriptions_url"`
		OrganizationsURL string `json:"organizations_url"`
		ReposURL string `json:"repos_url"`
		EventsURL string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type string `json:"type"`
		SiteAdmin bool `json:"site_admin"`
	} `json:"owner"`
	Private bool `json:"private"`
	HTMLURL string `json:"html_url"`
	Description string `json:"description"`
	Fork bool `json:"fork"`
	URL string `json:"url"`
	ForksURL string `json:"forks_url"`
	KeysURL string `json:"keys_url"`
	CollaboratorsURL string `json:"collaborators_url"`
	TeamsURL string `json:"teams_url"`
	HooksURL string `json:"hooks_url"`
	IssueEventsURL string `json:"issue_events_url"`
	EventsURL string `json:"events_url"`
	AssigneesURL string `json:"assignees_url"`
	BranchesURL string `json:"branches_url"`
	TagsURL string `json:"tags_url"`
	BlobsURL string `json:"blobs_url"`
	GitTagsURL string `json:"git_tags_url"`
	GitRefsURL string `json:"git_refs_url"`
	TreesURL string `json:"trees_url"`
	StatusesURL string `json:"statuses_url"`
	LanguagesURL string `json:"languages_url"`
	StargazersURL string `json:"stargazers_url"`
	ContributorsURL string `json:"contributors_url"`
	SubscribersURL string `json:"subscribers_url"`
	SubscriptionURL string `json:"subscription_url"`
	CommitsURL string `json:"commits_url"`
	GitCommitsURL string `json:"git_commits_url"`
	CommentsURL string `json:"comments_url"`
	IssueCommentURL string `json:"issue_comment_url"`
	ContentsURL string `json:"contents_url"`
	CompareURL string `json:"compare_url"`
	MergesURL string `json:"merges_url"`
	ArchiveURL string `json:"archive_url"`
	DownloadsURL string `json:"downloads_url"`
	IssuesURL string `json:"issues_url"`
	PullsURL string `json:"pulls_url"`
	MilestonesURL string `json:"milestones_url"`
	NotificationsURL string `json:"notifications_url"`
	LabelsURL string `json:"labels_url"`
	ReleasesURL string `json:"releases_url"`
	DeploymentsURL string `json:"deployments_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	PushedAt time.Time `json:"pushed_at"`
	GitURL string `json:"git_url"`
	SSHURL string `json:"ssh_url"`
	CloneURL string `json:"clone_url"`
	SvnURL string `json:"svn_url"`
	Homepage string `json:"homepage"`
	Size int `json:"size"`
	StargazersCount int `json:"stargazers_count"`
	WatchersCount int `json:"watchers_count"`
	Language string `json:"language"`
	HasIssues bool `json:"has_issues"`
	HasDownloads bool `json:"has_downloads"`
	HasWiki bool `json:"has_wiki"`
	HasPages bool `json:"has_pages"`
	ForksCount int `json:"forks_count"`
	MirrorURL interface{} `json:"mirror_url"`
	OpenIssuesCount int `json:"open_issues_count"`
	Forks int `json:"forks"`
	OpenIssues int `json:"open_issues"`
	Watchers int `json:"watchers"`
	DefaultBranch string `json:"default_branch"`
	Score float64 `json:"score"`
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
