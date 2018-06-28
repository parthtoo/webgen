package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"io/ioutil"
	"html/template"
	"gopkg.in/russross/blackfriday.v2"
)

const (
	PAGES_SRC = "src/pages/"
	PAGES_DST = "web/pages/"
	
	POSTS_SRC = "src/posts/"
	POSTS_DST = "web/posts/"
)

var (
	PAGE_TMPLS = template.Must(template.ParseFiles("src/fulls/page.html",
																"src/partials/head.html",
																"src/partials/nav.html"))
)

func mdToHtml(filename string) (string, error) {
	input, e := ioutil.ReadFile(filename)
	if e != nil { return "", e }
	output := blackfriday.Run(input, blackfriday.WithNoExtensions())
	return string(output), nil
}

func extractTitle(filename string) (string, error) {
	r, e := os.Open(filename)
	defer r.Close()
	if e != nil { return "", e }
	b := bufio.NewReader(r)
	tmp, e := b.ReadString('\n')
	if e != nil { return "", e }
	return tmp[:len(tmp) - 1], nil
}

func buildPages() {
	// read in pages
	infos, e := ioutil.ReadDir(PAGES_SRC)
	if e != nil {
		fmt.Printf("error reading in pages: %v\n", e)
		return
	}
	
	// call build for each
	for _, info := range infos {
		with_ext := info.Name()
		name := strings.TrimSuffix(with_ext, ".md")
		buildPage(name)
	}
}

func buildPosts() {
}

func buildPage(page string) {
	src_name := PAGES_SRC + page + ".md"
	dst_name := PAGES_DST + page + ".html"
	
	// confirm page exists
	if _, e := os.Stat(src_name); os.IsNotExist(e) {
		fmt.Printf("the page %v does not exist\n", page)
		return;
	}
	
	// convert markdown to html
	article, e := mdToHtml(src_name)
	if e != nil {
		fmt.Printf("error converting md to html for page %v: %v", page, e)
		return
	}
	
	// add article to templates
	tmpls, _ := PAGE_TMPLS.Clone()
	tmpls = tmpls.New("article.html")
	tmpls, e = tmpls.Parse(article)
	if e != nil {
		fmt.Printf("error parsing article for page %v: %v\n", page, e)
		return
	}
	
	// extract title
	title, e := extractTitle(src_name)
	if e != nil {
		fmt.Printf("error extracting title from page %v: %v", page, e)
	}
	
	// execute templates
	w, e := os.Create(dst_name)
	defer w.Close()
	if e != nil {
		fmt.Printf("error creating page %v: %v", page, e)
		return
	}
	e = tmpls.ExecuteTemplate(w, "page.html", title)
	if e != nil {
		fmt.Printf("error executing templates for page %v: %v\n", page, e)
		return
	}

	// at the end
	fmt.Printf("page %v successfully written\n", page)
}

func buildPost(post string) {
}