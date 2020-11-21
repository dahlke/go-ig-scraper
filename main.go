package main

import (
	"os"

	util "github.com/dahlke/goramma/util"
	worker "github.com/dahlke/goramma/worker"
)

func main() {
	util.ConfigLogger()

	instagramToken := os.Getenv("INSTAGRAM_ACCESS_TOKEN")
	worker.GetDataFromInstagramForUser(instagramToken)
}
