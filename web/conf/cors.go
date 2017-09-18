package conf

import (
	"net/http"
)

// 包装http.HandlerFunc,跨域
func CorsHandlerFunc(inner http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Method", "POST, OPTIONS, GET, HEAD, PUT, PATCH, DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "X-Requested-With, Content-Type")
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		inner.ServeHTTP(w, r)
	}
}
