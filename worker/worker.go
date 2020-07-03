package workers

import (
	"fmt"
	"time"

	api "github.com/dahlke/goramma/api"
	structs "github.com/dahlke/goramma/structs"
	log "github.com/sirupsen/logrus"
)

// TODO: remove this
const BaseUrl = "https://www.instagram.com"

/*
func writeInstagramMedia(filePath string, username string, media eklstructs.IGAPIMediaDetail) {
	fileContents, _ := json.MarshalIndent(media, "", " ")
	err := ioutil.WriteFile(filePath, fileContents, 0644)

	if err != nil {
		log.Error(err)
	} else {
		infoMsg := fmt.Sprintf("%s %s", filePath, "written")
		log.Info(infoMsg)
	}
}

func writeInstagramShortcodes(filePath string, media []eklstructs.IGAPITimelineMediaEdge, overwriteFiles bool) {
	fileContents := []byte{}
	if overwriteFiles {
		fileContents, _ = json.MarshalIndent(media, "", " ")
	} else {
		rawFileContents, _ := ioutil.ReadFile(filePath)
		existingMedia := []eklstructs.IGAPITimelineMediaEdge{}
		_ = json.Unmarshal([]byte(rawFileContents), &existingMedia)
		media = append(media, existingMedia...)
		fileContents, _ = json.MarshalIndent(media, "", " ")
	}

	err := ioutil.WriteFile(filePath, fileContents, 0644)
	if err != nil {
		log.Error(err)
	} else {
		log.Info("Instagram shortcodes written")
	}
}

func writeInstagramEndCursor(endCursor string) {
	fileWritePath := fmt.Sprintf("./data/instagram/worker/latest-consumed-end-cursor.json")
	fileWriteAbsPath, err := filepath.Abs(fileWritePath)
	if err != nil {
		log.Error(err)
	}

	fileContents := []byte(endCursor)
	err = ioutil.WriteFile(fileWriteAbsPath, fileContents, 0644)

	if err != nil {
		log.Error(err)
	} else {
		infoMsg := fmt.Sprintf("Latest end cursor saved.")
		log.Info(infoMsg)
	}
}

func writeInstagramSimpleMedia(simpleMedia []eklstructs.InstagramMedia) {
	fileWritePath := fmt.Sprintf("./data/instagram-simplified.json")
	fileWriteAbsPath, err := filepath.Abs(fileWritePath)
	if err != nil {
		log.Error(err)
	}

	fileContents, _ := json.MarshalIndent(simpleMedia, "", " ")
	err = ioutil.WriteFile(fileWriteAbsPath, fileContents, 0644)

	if err != nil {
		log.Error(err)
	} else {
		infoMsg := fmt.Sprintf("Instagram simplified data written")
		log.Info(infoMsg)
	}

}

func consolidateAndConvertMedia(directory string) {
	files, err := ioutil.ReadDir(directory)

	if err != nil {
		log.Fatal(err)
	}

	allMedia := []eklstructs.IGAPIMediaDetail{}
	for _, f := range files {
		filePath := fmt.Sprintf("%s/%s", directory, f.Name())
		rawFileContents, _ := ioutil.ReadFile(filePath)
		var media eklstructs.IGAPIMediaDetail
		_ = json.Unmarshal([]byte(rawFileContents), &media)
		allMedia = append(allMedia, media)
	}

	allSimpleMedia := []eklstructs.InstagramMedia{}
	for _, m := range allMedia {
		simpleMedia := eklstructs.InstagramMedia{
			m.GraphQL.Media.Shortcode,
			m.GraphQL.Media.Timestamp,
			m.GraphQL.Media.Location.Name,
			m.GraphQL.Media.DisplayURL,
		}
		allSimpleMedia = append(allSimpleMedia, simpleMedia)
	}

	writeInstagramSimpleMedia(allSimpleMedia)
}
*/

func GetDataFromInstagramForUser(username string, endCursor string) {
	// TODO: take an encursor instead of pull All Instagrams

	// Retrieve user metadata
	// TODO: move this to a setup function so we don't call it here but we don't have to pass it all subsequent functions?
	userID := api.GetUserIDFromMetadata(username)
	pullAll := endCursor != ""
	var timelineEdges []structs.IGAPITimelineMediaEdge

	// Query or the Instagram media
	for true {
		timeline := api.GetUserTimelineMedia(userID, endCursor)
		timelineEdges = append(timeline.Data.User.Media.Edges, timelineEdges...)
		hasNextPage := timeline.Data.User.Media.PageInfo.HasNextPage

		endCursor = timeline.Data.User.Media.PageInfo.EndCursor

		if hasNextPage && pullAll {
			log.Info("Getting another page...")
		} else {
			fmt.Println("breaking")
			break
		}

		time.Sleep(3 * time.Second)
	}
	fmt.Println(timelineEdges)
	/*
		writeInstagramShortcodes(fileAbsPath, timelineEdges, pullAllInstagrams)

		mediaDir := "./data/instagram/worker/media"
		for i, edge := range timelineEdges {
			shortcode := edge.Node.Shortcode
			filePath := fmt.Sprintf("%s/%s.json", mediaDir, shortcode)
			fileAbsPath, err := filepath.Abs(filePath)

			if err != nil {
				log.Error(err)
			}

			if !fileExists(fileAbsPath) {
				log.Info(fmt.Sprintf("Getting the details for media %d of %d...", i, len(timelineEdges)))
				mediaDetail := getMediaDetailFromShortcode(shortcode)
				writeInstagramMedia(fileAbsPath, username, *mediaDetail)
			}
		}

		consolidateAndConvertMedia(mediaDir)
	*/

	// TODO: output last endcursor
}
