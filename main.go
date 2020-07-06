package main

import (
	"flag"

	util "github.com/dahlke/goramma/util"
	worker "github.com/dahlke/goramma/worker"
)

func main() {
	usernamePtr := flag.String("-username", "eklhad", "Username to scrape.")
	endCursorPtr := flag.String("-endcursor", "", "The endcursor to mark where to start pulling data from, otherwise starts at the beginning of the timeline.")
	flag.Parse()

	util.ConfigLogger()
	username := *usernamePtr
	endCursor := *endCursorPtr

	worker.GetDataFromInstagramForUser(username, endCursor)
}
