package conf

import (
	"time"
	"log"
	"net/http"
)

// 包装http.HandlerFunc,打印日志信息
func LoggerHandlerFunc(inner http.HandlerFunc, handlerFuncName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf("%s\t%s\t%s\t%s.\n", r.Method, r.RequestURI, handlerFuncName, time.Since(start))
	}
}