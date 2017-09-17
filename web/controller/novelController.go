package controller

import (
	"net/http"
	"encoding/json"
	"github.com/xuanbo/novel/web/conf"
	"github.com/gorilla/mux"
	"github.com/xuanbo/novel/core/source"
	"github.com/xuanbo/novel/core/model"
	"strings"
)

var UserController = &conf.Controller{
	Name: "novelController",
	Path: "/novel",
	Routes: &[]conf.Route {
		{
			Match: "/search/{q}",
			Methods: &[]string{conf.GET},
			HandleFunc: search,
		},
		{
			Match: "",
			Methods: &[]string{conf.GET},
			HandleFunc: getNovel,
		},
		{
			Match: "/chapter",
			Methods: &[]string{conf.GET},
			HandleFunc: getChapter,
		},
	},
}

func search(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	q := vars["q"]
	result, err := source.BiqugeFetcher.Search(q)
	if err != nil {
		writer.Write(toJsonByte(model.Fail("")))
		return
	}
	writer.Write(toJsonByte(model.Ok(result)))
}

func getNovel(writer http.ResponseWriter, req *http.Request) {
	url := req.URL.Query().Get("url")
	if strings.Trim(url, "") == "" {
		writer.Write(toJsonByte(model.Fail("参数不正确")))
		return
	}
	result, err := source.BiqugeFetcher.FetchNovel(url)
	if err != nil {
		writer.Write(toJsonByte(model.Fail("")))
		return
	}
	writer.Write(toJsonByte(model.Ok(result)))
}

func getChapter(writer http.ResponseWriter, req *http.Request) {
	url := req.URL.Query().Get("url")
	if strings.Trim(url, "") == "" {
		writer.Write(toJsonByte(model.Fail("参数不正确")))
		return
	}
	result, err := source.BiqugeFetcher.FetchChapter(url)
	if err != nil {
		writer.Write(toJsonByte(model.Fail("")))
		return
	}
	writer.Write(toJsonByte(model.Ok(result)))
}

func toJsonByte(data interface{}) []byte {
	result, err := json.Marshal(data)
	if err != nil {
		return []byte("")
	}
	return result
}