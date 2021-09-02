package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page, Hello GoLang.<a href=\"/about\">About</a>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "about page 关于我们 <a href=\"/\">Home</a>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "404 not found page")
}

func articlesShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprint(w, "文章ID："+id)
}

func articlesIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "访问文章列表")
}

func articlesStoreHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprint(w, "请提交正确的数据！")
		return
	}
	title := r.PostForm.Get("title")

	fmt.Fprintf(w, "r.Form 中 title 的值为: %v <br>", r.FormValue("title"))
	fmt.Fprintf(w, "r.PostForm 中 title 的值为: %v <br>", r.PostFormValue("title"))
	fmt.Fprintf(w, "Post PostForm: %v <br />", r.PostForm)
	fmt.Fprintf(w, "Post Form: %v <br />", r.Form)
	fmt.Fprintf(w, "title: %v", title)
}

func articlesCreateHandler(w http.ResponseWriter, r *http.Request) {
	html := `
	<!doctype html>
	<html lang="en">
	<head>
	<title>创建文章</title>
	</head>
	<body>
	      <form action="%s?test=data" method="post">
		     <p><input type="text" name="title"></p>
			 <p><textarea name="body" cols="30" rows="10"></textarea></p>
			 <p><button type="submit">提交</button></p>
		  </form>
	</body>
	</html>
	`
	storeUrl, _ := router.Get("articles.store").URL()
	fmt.Fprintf(w, html, storeUrl)
}

func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		next.ServeHTTP(w, r)
	})
}

func main() {

	router.HandleFunc("/", homeHandler).Methods("GET").Name("home")
	router.HandleFunc("/about", aboutHandler).Methods("GET").Name("about")

	router.HandleFunc("/articles/{id:[0-9]+}", articlesShowHandler).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles", articlesIndexHandler).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles", articlesStoreHandler).Methods("POST").Name("articles.store")
	router.HandleFunc("/articles/create", articlesCreateHandler).Methods("GET").Name("articles.create")

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	// 中间件
	router.Use(forceHTMLMiddleware)

	http.ListenAndServe(":3000", removeTrailingSlash(router))
}
