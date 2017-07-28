package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gyukebox/gyulog/post"
)

// temporary
func index(w http.ResponseWriter, r *http.Request) {
	var beginindex int
	end, err := url.QueryUnescape(r.URL.RawQuery)
	if err != nil {
		beginindex = 0
	} else {
		beginindex, _ = strconv.Atoi(end)
	}
	if err != nil {
		fmt.Print("At Handler : ")
		fmt.Println(err)
	}
	posts, err := post.GetFivePosts(beginindex)
	generateHTML(w, posts, "index", "layout", "mobile", "navbar", "pager", "sidebar")
}

func postDetail(w http.ResponseWriter, r *http.Request) {
	title, err := url.QueryUnescape(r.URL.RawQuery)
	if err != nil {
		fmt.Print("At parsing url : ")
		fmt.Println(err)
		log.Fatalln(err)
	}
	post := post.GetPostByTitle(title)
	data := map[string]template.HTML{
		"Title":         template.HTML(post.Title),
		"Body":          template.HTML(post.Body),
		"PublishedDate": template.HTML(post.PublishedDate),
		"Summary":       template.HTML(post.Summary),
	}
	generateHTML(w, data, "post", "layout", "mobile", "navbar", "sidebar", "pager")
}
