package middlewares

import "net/http"

func ForceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 设置标头
		//w.Header().Set("Content-Type", "text/css; charset=utf-8")
		// 2. 继续处理请求
		next.ServeHTTP(w, r)
	})
}
