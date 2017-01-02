package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHumanReadable(t *testing.T) {
	assert.Equal(t, "1.15 Gb", HumanReadableBytes(1234803097))
	assert.Equal(t, "3.26 Mb", HumanReadableBytes(3418357))
	assert.Equal(t, "7.64 kb", HumanReadableBytes(7823))
	assert.Equal(t, "823 b", HumanReadableBytes(823))
}

func TestSearchAPI(t *testing.T) {
	search, _ := SearchGithub("golang", 32)
	assert.Equal(t, 32, len(search))
}

func TestLanguagesAPI(t *testing.T) {
	repo := GithubRepo{LanguagesURL: "https://api.github.com/repos/kephas/jedossi/languages"}
	repo.GetLanguagesData()
	assert.True(t, repo.Languages["Elixir"] > 0)
}
