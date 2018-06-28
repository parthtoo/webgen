package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"github.com/parthtoo/webgen/pnp"
)

func clean() {}
func buildPosts() {}

func buildPages() {
	infos, e := ioutil.ReadDir(WEBSITE + "/src/pages")
	if e != nil {
		fmt.Printf("error reading in pages from directory:\n%v\n", e)
		return
	}
	for _, info := range infos {
		tmp := info.Name()
		name := strings.TrimSuffix(tmp, ".md")
		buildPage(name)
	}
}

func buildPage(name string) {
	p, e := pnp.InitPage(name)
	if e != nil {
		fmt.Printf("error while initializing page %v:\n%v\n", name, e)
		return
	}
	if e = p.Load(); e != nil {
		fmt.Printf("error while loading page %v:\n%v\n", name, e)
		return
	}
	if e = p.Write(); e != nil {
		fmt.Printf("error while writing page %v:\n%v\n", name, e)
		return
	}
	fmt.Printf("page %v successfully written\n", name)
}

func buildPost(name string) {
	// p, e := InitPost(name)
	// if e != nil {
	// 	fmt.Printf("error encountered while initializing post %v: %v\n", name, e)
	// 	return
	// }
	// e = p.LoadMeta()
	// if e != nil {
	// 	fmt.Printf("error encountered while loading metadata of post %v: %v\n", name, e)
	// 	return
	// }
	// e = p.LoadCntnt()
	// if e != nil {
	// 	fmt.Printf("error encountered while loading content of post %v: %v\n", name, e)
	// 	return
	// }
	// e = p.Build()
	// if e != nil {
	// 	fmt.Printf("error encountered while building post %v: %v\n", name, e)
	// 	return
	// }
}