package model

// 定义章节
type Chapter struct {
	Name string `json:"name"` // 章节名称
	Url string `json:"url"` // 章节url
	Content string `json:"content"` // 章节内容
}