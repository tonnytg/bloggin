package main

import "flag"

func main() {

	// create prompt to manage blog

	// add flag --add post
	// add flag --list post
	// add flag --delete post

	action := flag.String("action", "list", "--action list")
	flag.Parse()

	if *action == "list" {
		// list all posts
	} else if *action == "add" {
		// add a post
	} else if *action == "delete" {
		// delete a post
	} else {
		// show help
	}
}
