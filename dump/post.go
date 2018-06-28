package main

import (
	"os"
	"fmt"
	"bytes"
	"errors"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"gopkg.in/russross/blackfriday.v2"
)

type Post struct {
	name string
	src_file string
	dst_file string
	
	meta map[interface{}]interface{}
	cntnt string
}

type Related map[string]string

func InitPost(name string) (*Post, error) {
	src_file := WEBSITE + "/src/posts/" + name + ".md"
	if _, e := os.Stat(src_file); os.IsNotExist(e) {
		return nil, errors.New(fmt.Sprintf("the post %v does not exist\n", name))
	}
	dst_file := WEBSITE + "/web/posts/" + name + ".html"
	return &Post{name, src_file, dst_file, nil, ""}, nil
}

func (p *Post) LoadMetaShallow() error {
	all, e := ioutil.ReadFile(p.src_file)
	if e != nil { return e }
	sep := bytes.Index(all, []byte("..."))
	if sep == -1 {
		return errors.New(fmt.Sprintf("no separator ... found in post %v\n", p.name))
	}
	meta := all[:sep]
	p.meta = make(map[interface{}]interface{})
	e = yaml.Unmarshal(meta, p.meta)
	if e != nil { return e }
	return nil
}

func (p *Post) LoadMeta() error {
	e := p.LoadMetaShallow()
	if e != nil { return e }
	
	// related posts
	rel := p.meta["related"].([]interface{})
	p.meta["related-new"] = make([]Related, len(rel))
	for i, tmp := range rel {
		r := tmp.(string)
		pr, e := InitPost(r)
		if e != nil { return e }
		e = pr.LoadMetaShallow()
		if e != nil { return e }
		
		new_rel := p.meta["related-new"].([]Related)
		new_rel[i] = make(Related)
		new_rel[i]["title"] = pr.meta["title"].(string)
		new_rel[i]["url"] = "/posts/" + pr.name + ".html"
	}
	
	fmt.Println(p.meta)
	return nil
}

func (p *Post) LoadCntnt() error {
	all, e := ioutil.ReadFile(p.src_file)
	if e != nil { return e }
	sep := bytes.Index(all, []byte("..."))
	if sep == -1 {
		return errors.New(fmt.Sprintf("no separator ... found in post %v\n", p.name))
	}
	md := all[sep + 3:]
	html := blackfriday.Run(md, blackfriday.WithNoExtensions())
	p.cntnt = string(html)
	
	fmt.Println(p.cntnt)
	return nil
}