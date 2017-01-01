package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestHumanReadable(t *testing.T) {
	assert.Equal(t, "1.15 Gb", HumanReadableBytes(math.Pow(1024, 3) * 1.15))
	assert.Equal(t, "3.26 Mb", HumanReadableBytes(math.Pow(1024, 2) * 3.26))
	assert.Equal(t, "7.64 kb", HumanReadableBytes(1024 * 7.64))
	assert.Equal(t, "823 b", HumanReadableBytes(823))
}

func TestSearchAPI(t *testing.T) {
	search, _ := SearchGithub("golang", 32)
	assert.Equal(t, 32, len(search))
}
