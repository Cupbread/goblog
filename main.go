package main

import (
	"fmt"
	"net/http"
)

func defaultFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello, 部署热更新！</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>"+
			"<p>如有疑惑，请联系我们。</p>")
	}
}

func aboutFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:shucanwu@163.com\">shucanwu@163.com</a>")
}

func contactFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "123")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", defaultFunc)
	router.HandleFunc("/about", aboutFunc)
	router.HandleFunc("/contact", contactFunc)

	http.ListenAndServe(":3000", router)
}
