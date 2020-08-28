package workers

import (
	"time"

	api "github.com/dahlke/goramma/api"
	structs "github.com/dahlke/goramma/structs"
	log "github.com/sirupsen/logrus"
)

// GetDataFromInstagramForUser is used for testing that the API functions work as expected.
func GetDataFromInstagramForUser(username string, inputEndCursor string) {
	userID := api.GetUserIDFromMetadata(username)
	pullAll := inputEndCursor == ""
	var fullConvertedMediaTimeline []structs.InstagramMedia
	endCursor := inputEndCursor

	for true {
		mediaTimeline, hasNextPage, newEndCursor := api.GetUserTimelineMedia(userID, endCursor)
		endCursor = newEndCursor
		fullConvertedMediaTimeline = append(mediaTimeline, fullConvertedMediaTimeline...)

		if hasNextPage && pullAll {
			log.Info("Getting another page...")
		} else {
			break
		}

		time.Sleep(3 * time.Second)
	}
	// TODO: log the end cursor.
	// fmt.Println(endCursor, fullConvertedMediaTimeline)
}
