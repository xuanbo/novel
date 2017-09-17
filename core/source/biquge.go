package source

import (
	"net/url"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"github.com/xuanbo/novel/core/model"
)

// 笔趣阁
type Biquge struct {
}

const (
	SOURCE_TYPE = 1
	SEARCH_URL = "http://www.biquge5200.com/modules/article/search.php?searchkey="
)

var BiqugeFetcher = &Biquge{}

// 实现 source.Fetcher 接口
func (Biquge) Search(word string) (*[]model.Search, error) {
	url := SEARCH_URL + url.PathEscape(word)
	document, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	return parseSearch(document), nil
}

func (Biquge) FetchNovel(url string) (*model.Novel, error) {
	document, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	return parseNovel(document), nil
}

func (Biquge) FetchChapter(url string) (*model.Chapter, error) {
	document, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	return parseChapter(document)
}

// 解析小说搜索结果
func parseSearch(doc *goquery.Document) *[]model.Search {
	result := make([]model.Search, 0)
	doc.Find("#hotcontent>.grid>tbody>tr").Each(func(i int, selection *goquery.Selection) {
		if i == 0 {
			return
		}
		result = append(result, parseSearchItem(selection))
	})
	return &result
}

// 解析搜索结果条目
func parseSearchItem(selection *goquery.Selection) model.Search {
	search := model.Search{Source: SOURCE_TYPE}
	selection.Find("td").Each(func(i int, selection *goquery.Selection) {
		switch i {
			case 0:
				a := selection.Find("a")
				novelUrl, exists := a.Attr("href")
				if exists {
					search.NovelUrl = toUtf8(novelUrl)
				}
				search.NovelName = toUtf8(a.Text())
			case 1:
				a := selection.Find("a")
				search.LastChapterName = toUtf8(a.Text())
			case 2:
				search.Author = toUtf8(selection.Text())
			case 3:
				search.WordCount = toUtf8(selection.Text())
			case 4:
				search.UpdateTime = toUtf8(selection.Text())
			case 5:
				search.Status = toUtf8(selection.Text())
		}
	})
	return search
}


// 解析小说
func parseNovel(doc *goquery.Document) *model.Novel {
	name := parseNovelName(doc)
	author := parseNovelAuthor(doc)
	lastUpdateTime := parseNovelLastUpdateTime(doc)
	description := parseNovelDescription(doc)
	novel := &model.Novel{
		Name: 	name,
		Author: author,
		LastUpdateTime: lastUpdateTime,
		Description: description,
		Chapters: parseNovelChapters(doc),
		Source: 1,
	}
	return novel
}

// 小说名
func parseNovelName(doc *goquery.Document) string {
	name:= doc.Find("#info>h1").Text()
	return toUtf8(name)
}

// 小说作者
func parseNovelAuthor(doc *goquery.Document) string {
	author := doc.Find("#info>p").First().Text()
	author = toUtf8(author)
	author = strings.Replace(author, "作聽聽聽聽者：", "", -1)
	return author
}

// 小说最后一次更新时间
func parseNovelLastUpdateTime(doc *goquery.Document) string {
	lastUpdateTime := doc.Find("#info>p").Last().Text()
	lastUpdateTime = toUtf8(lastUpdateTime)
	lastUpdateTime = strings.Replace(lastUpdateTime, "最后更新：", "", -1)
	return lastUpdateTime
}

// 小说描述
func parseNovelDescription(doc *goquery.Document) string {
	description := doc.Find("#intro>p").Last().Text()
	return toUtf8(description)
}

// 小说的章节
func parseNovelChapters(doc *goquery.Document) *[]model.Chapter {
	chapters := make([]model.Chapter, 0)
	doc.Find("#list>dl>dd").Each(func(i int, selection *goquery.Selection) {
		if i < 9 {
			return
		}
		a := selection.Find("a")
		url, exists := a.Attr("href")
		url = toUtf8(url)
		if !exists {
			return
		}
		name := a.Text()
		name = toUtf8(name)
		chapters = append(chapters, model.Chapter{Name: name, Url: url})
	})
	return &chapters
}

// 解析章节
func parseChapter(doc *goquery.Document) (*model.Chapter, error) {
	name := doc.Find(".bookname>h1").Text()
	content, err := doc.Find("#content").Html()
	if err != nil {
		return nil, err
	}
	name = toUtf8(name)
	content = toUtf8(content)
	content = strings.Replace(content, "<br/>", "\n\n", -1)
	return &model.Chapter{Name: name, Content: content}, nil
}

func toUtf8(text string) string {
	return mahonia.NewDecoder("gbk").ConvertString(text)
}