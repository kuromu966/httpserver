package tpl_error

import (
	"encoding/json"
	"io/ioutil"
)

type PageContent struct {
	Title string `json:"title"`
}

func LoadContent(path string) *PageContent {
	var result PageContent
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(file, &result)

	return &result
}
