package workers

import (
	"fmt"
	"time"

	api "github.com/dahlke/goramma/api"
	"github.com/dahlke/goramma/structs"

	// structs "github.com/dahlke/goramma/structs"
	log "github.com/sirupsen/logrus"
)

// GetDataFromInstagramForUser is used for testing that the API functions work as expected.
func GetDataFromInstagramForUser(instagramToken string) {
	userMetadata := api.GetUserMetadata(instagramToken)
	fmt.Println(userMetadata)

	var allInstagramMedia []structs.InstagramMedia
	beforeEndCursor := ""
	endCursor := ""

	for true {
		log.Info(fmt.Sprintf("Fetching page of Instagram data with end cursor: %s ...", endCursor))
		userMedia := api.GetUserMedia(instagramToken, endCursor)
		allInstagramMedia = append(allInstagramMedia, userMedia.Data...)
		beforeEndCursor = userMedia.Paging.Cursors.Before
		endCursor = userMedia.Paging.Cursors.After

		if beforeEndCursor == endCursor {
			log.Info("Finished fetching data from Instagram.")
			break
		}
		time.Sleep(1 * time.Second)
	}

	fmt.Println(allInstagramMedia)
}
