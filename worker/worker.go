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

func GetDataFromInstagramForUser(username string, endCursor string) {
	// TODO: take an encursor instead of pull All Instagrams

	// Retrieve user metadata
	// TODO: move this to a setup function so we don't call it here but we don't have to pass it all subsequent functions?
	userID := api.GetUserIDFromMetadata(username)
	pullAll := endCursor != ""
	var fullMediaTimeline []structs.IGAPIMediaDetail

	// Query or the Instagram media
	for true {
		mediaTimeline, hasNextPage, endCursor := api.GetUserTimelineMedia(userID, endCursor)
		fullMediaTimeline = append(mediaTimeline, fullMediaTimeline...)
		fmt.Println(mediaTimeline, endCursor, hasNextPage)

		if hasNextPage && pullAll {
			log.Info("Getting another page...")
		} else {
			fmt.Println("breaking")
			break
		}

		time.Sleep(3 * time.Second)
	}

	// TODO: output last endcursor
	fmt.Println(fullMediaTimeline, endCursor)

}
