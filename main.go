package main

import (
	"flag"
	"fmt"

	util "github.com/dahlke/goramma/util"
	worker "github.com/dahlke/goramma/worker"
)

func main() {
	// TODO: config path
	// TODO: take an endcursor argument
	configPathPtr := flag.String("-config-path", "config.example.json", "Path to config file, example included as `config.example.json`")
	endCursorPtr := flag.String("-endcursor", "", "The endcursor to mark where to start pulling data from, otherwise starts at the beginning of the timeline.")
	// metadataPathPtr := flag.String("-metadata-path", "", "The path to write all the metadata about your images to.")
	// imagesPathPtr := flag.String("-images-path", "", "The path to write all the images to.")
	flag.Parse()

	util.ConfigLogger()
	configPath := *configPathPtr
	endCursor := *endCursorPtr
	// metadataPath := *metadataPathPtr
	// imagesPath := imagesPathPtr

	appConfig := util.ParseConfig(configPath)
	fmt.Println(configPath, endCursor, appConfig)

	worker.GetDataFromInstagramForUser(appConfig.InstagramUsername, endCursor)
	// TODO: remember that this library is meant to be imported and not the actual worker, but keep the worker for testing.
}
