package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func defaultFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, 默认页面！</h1>")
}

func aboutFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:shucanwu@163.com\">shucanwu@163.com</a>")
}

func articlesShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprint(w, "文章 ID："+id)
}

func articlesIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "访问文章列表")
}

func articlesStoreHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "创建新的文章")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
}

func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", defaultFunc).Methods("GET").Name("default")
	router.HandleFunc("/about", aboutFunc).Methods("GET").Name("about")

	router.HandleFunc("/articles/{id:[0-9]+}", articlesShowHandler).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles", articlesIndexHandler).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles", articlesStoreHandler).Methods("GET").Name("articles.store")

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	router.Use(forceHTMLMiddleware)

	http.ListenAndServe(":3000", router)
}
