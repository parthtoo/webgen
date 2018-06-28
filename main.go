package main

import (
	"os"
	"fmt"
	"log"
)

var (
	WEBSITE = os.Getenv("WEBSITE")
)

func confirmWebsite() {
	if WEBSITE == "" {
		log.Fatal(fmt.Errorf("please provide location of the website by setting environment variable WEBSITE"))
	}
	_, e1 := os.Stat(WEBSITE + "/src")
	_, e2 := os.Stat(WEBSITE + "/web")
	if os.IsNotExist(e1) || os.IsNotExist(e2) {
		log.Fatal(fmt.Errorf("the provided website does not contain src and/or web directories"))
	}
}

func routeCommand() {
	// error checking
	if len(os.Args) == 1 {
		log.Fatal(fmt.Errorf("please provide a command"))
	}

	switch cmd := os.Args[1]; cmd {
	case "build":
		// error checking
		if len(os.Args) == 2 {
			log.Fatal(fmt.Errorf("please specify what to build"))
		}

		switch sub_cmd := os.Args[2]; sub_cmd {
		case "all" :
			if len(os.Args) == 3 {
				buildPages()
				buildPosts()
				return
			}
			switch sub_cmd_two := os.Args[3]; sub_cmd_two {
			case "pages":
				buildPages()
			case "posts":
				buildPosts()
			default:
				log.Fatal(fmt.Errorf("cannot build all %v", sub_cmd_two))
			}
		
		case "pages":
			// error checking
			if len(os.Args) == 3 {
				log.Fatal(fmt.Errorf("please provide pages to build"))
			}
			ps := os.Args[3:]
			for _, p := range ps {
				buildPage(p)
			}
		
		case "posts":
			// error checking
			if len(os.Args) == 3 {
				log.Fatal(fmt.Errorf("please provide posts to build"))
			}
			ps := os.Args[3:]
			for _, p := range ps {
				buildPost(p)
			}
		
		default:
			log.Fatal(fmt.Errorf("cannot build %v", sub_cmd))
		}
	case "clean":
		clean()
	default:
		log.Fatal(fmt.Errorf("please provide a valid command"))
	}
}

func main() {
	confirmWebsite()
	routeCommand()
}