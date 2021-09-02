package main

import (
	"fmt"
	"net/http"
)

func defaultFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, "Home Page, Hello GoLang.<a href=\"/about\">About</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 错误页")
	}
}

func aboutFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "about page 关于我们 <a href=\"/\">Home</a>")
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", defaultFunc)
	router.HandleFunc("/about", aboutFunc)

	http.ListenAndServe(":3000", router)
}
