package main

import (
	util "github.com/dahlke/goramma/util"
	worker "github.com/dahlke/goramma/worker"
)

func main() {
	util.ConfigLogger()

	worker.GetDataFromInstagramForUser()
}
