package model

// 定义小说
type Novel struct {
	Name string `json:"name"` // 小说名
	Author string `json:"author"` // 作者名
	LastUpdateTime string `json:"lastUpdateTime"` // 最后一次更新时间
	Description string `json:"description"` // 小说描述
	Chapters *[]Chapter `json:"chapters"`  // 章节列表 Chapter
	Source int `json:"source"` // 定义小说来源
}