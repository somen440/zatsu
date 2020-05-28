package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReflectTag(t *testing.T) {
	type Data struct {
		ID string `json:"id" path:"id"`
	}
	var data Data

	list := map[string]string{
		"id":   "aaa",
		"name": "bbb",
	}
	assert.Nil(t, ReflectTag("path", func(tag string) string { return list[tag] }, &data))
	assert.Equal(t, "aaa", data.ID)
}
