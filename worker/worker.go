package workers

import (
	"time"

	api "github.com/dahlke/goramma/api"
	structs "github.com/dahlke/goramma/structs"
	log "github.com/sirupsen/logrus"
)

// GetDataFromInstagramForUser is used for testing that the API functions work as expected.
func GetDataFromInstagramForUser(username string, endCursor string) {
	userID := api.GetUserIDFromMetadata(username)
	pullAll := endCursor != ""
	var fullConvertedMediaTimeline []structs.InstagramMedia

	for true {
		// mediaTimeline, hasNextPage, endCursor := api.GetUserTimelineMedia(userID, endCursor)
		mediaTimeline, hasNextPage, _ := api.GetUserTimelineMedia(userID, endCursor)
		fullConvertedMediaTimeline = append(mediaTimeline, fullConvertedMediaTimeline...)

		if hasNextPage && pullAll {
			log.Info("Getting another page...")
		} else {
			break
		}

		time.Sleep(3 * time.Second)
	}
}
