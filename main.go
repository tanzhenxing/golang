package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, "<h1>Hello GoLang Go 语言</h1>")
	} else if r.URL.Path == "/about" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, "請求路徑：", r.URL.Path)
		fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
			"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
	} else {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, "無法找到該頁面")
		fmt.Fprint(w, "<a href=\"/about\" >关于我们</a>", r.URL.Query())
		fmt.Fprint(w, r.Header.Get("User-Agent"))
	}

}
func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
