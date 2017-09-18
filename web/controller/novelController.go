package controller

import (
	"net/http"
	"strings"
	"github.com/xuanbo/novel/web/conf"
	"github.com/gorilla/mux"
	"github.com/xuanbo/novel/core/source"
	"github.com/xuanbo/novel/core/model"
	"github.com/xuanbo/novel/web/util"
)

var UserController = &conf.Controller {
	Name: "novelController",
	Path: "/novel",
	Routes: &conf.Routes {
		{
			Match: "/search/{q}",
			Methods: []string{conf.GET},
			HandleFuncName: "search",
			HandleFunc: search,
		},
		{
			Match: "",
			Methods: []string{conf.GET},
			HandleFuncName: "getNovel",
			HandleFunc: getNovel,
		},
		{
			Match: "/chapter",
			Methods: []string{conf.GET},
			HandleFuncName: "getChapter",
			HandleFunc: getChapter,
		},
	},
}

// 关键词检索小说
func search(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	q := vars["q"]
	result, err := source.BiqugeFetcher.Search(q)
	sendResult(result, err, writer)
}

// 获取小说信息
func getNovel(writer http.ResponseWriter, req *http.Request) {
	url := req.URL.Query().Get("url")
	if strings.Trim(url, "") == "" {
		writer.Write(util.ToJsonByte(model.Fail("参数不正确")))
		return
	}
	result, err := source.BiqugeFetcher.FetchNovel(url)
	sendResult(result, err, writer)
}

// 获取小说章节信息
func getChapter(writer http.ResponseWriter, req *http.Request) {
	url := req.URL.Query().Get("url")
	if strings.Trim(url, "") == "" {
		writer.Write(util.ToJsonByte(model.Fail("参数不正确")))
		return
	}
	result, err := source.BiqugeFetcher.FetchChapter(url)
	sendResult(result, err, writer)
}

func sendResult(result interface{}, err error, writer http.ResponseWriter) {
	if err != nil {
		writer.Write(util.ToJsonByte(model.Fail("")))
		return
	}
	writer.Write(util.ToJsonByte(model.Ok(result)))
}