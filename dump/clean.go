package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func clean() {
	// pages
	infos, _ := ioutil.ReadDir(PAGES_DST)
	for _, info := range infos {
		os.Remove(PAGES_DST + info.Name())
	}
	if len(infos) == 0 {
		fmt.Printf("no pages to delete\n")
	} else {
		fmt.Printf("all pages deleted\n")
	}
	
	// posts
	infos, _ = ioutil.ReadDir(POSTS_DST)
	for _, info := range infos {
		os.Remove(POSTS_DST + info.Name())
	}
	if len(infos) == 0 {
		fmt.Printf("no posts to delete\n")
	} else {
		fmt.Printf("all posts deleted\n")
	}
}