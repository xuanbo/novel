package model

type Search struct {
	NovelName string `json:"novelName"`
	NovelUrl string `json:"novelUrl"`
	LastChapterName string `json:"lastChapterName"`
	Author string `json:"author"`
	WordCount string `json:"wordCount"`
	UpdateTime string `json:"updateTime"`
	Status string `json:"status"`
	Source int `json:"source"`
}