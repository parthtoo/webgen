package pnp

import (
	"os"
	"fmt"
	"bytes"
	"io/ioutil"
	"html/template"
	"gopkg.in/russross/blackfriday.v2"
)

var (
	WEBSITE = os.Getenv("WEBSITE")
	TMPLS = template.Must(template.ParseFiles(WEBSITE + "/src/fulls/page.html",
															WEBSITE + "/src/partials/head.html",
															WEBSITE + "/src/partials/nav.html"))
)

type Page struct {
	Name string
	src_file string
	dst_file string
	
	title string
	cntnt string
}

func InitPage(name string) (*Page, error) {
	src_file := WEBSITE + "/src/pages/" + name + ".md"
	if _, e := os.Stat(src_file); os.IsNotExist(e) {
		return nil, fmt.Errorf("the page %v does not exist", name)
	}
	dst_file := WEBSITE + "/web/pages/" + name + ".html"
	return &Page{name, src_file, dst_file, "", ""}, nil
}

func (p *Page) IsLoaded() bool {
	return p.title != "" && p.cntnt != ""
}

func (p *Page) Load() error {
	if p.IsLoaded() {
		return fmt.Errorf("page %v is already loaded", p.Name)
	}
	md, e := ioutil.ReadFile(p.src_file)
	if e != nil {
		return fmt.Errorf("error while reading page %v\n%v", p.Name, e)
	}
	sep := bytes.Index(md, []byte("\n"))
	if sep == -1 {
		return fmt.Errorf("could not load title from page %v", p.Name)
	}
	p.title = string(md[:sep])
	html := blackfriday.Run(md, blackfriday.WithNoExtensions())
	p.cntnt = string(html)
	return nil
}

func (p *Page) Write() error {
	if !p.IsLoaded() {
		return fmt.Errorf("page %v is not loaded", p.Name)
	}
	tmpls, e := TMPLS.Clone()
	if e != nil {
		return fmt.Errorf("error while cloning TEMPLATES\n%v", e)
	}
	tmpls = tmpls.New("cntnt.html")
	tmpls, e = tmpls.Parse(p.cntnt)
	if e != nil {
		return fmt.Errorf("error while parsing content of page %v\n%v", p.Name, e)
	}
	w, e := os.Create(p.dst_file)
	if e != nil {
		return fmt.Errorf("error opening file for writing for page %v\n%v", p.Name, e)
	}
	defer w.Close()
	e = tmpls.ExecuteTemplate(w, "page.html", p.title)
	if e != nil {
		return fmt.Errorf("error while executing templates of page %v\n%v", p.Name, e)
	}
	return nil
}
