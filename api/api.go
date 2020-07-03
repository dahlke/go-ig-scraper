package api

import (
	"encoding/json"
	"fmt"

	"github.com/dahlke/goramma/structs"
	log "github.com/sirupsen/logrus"
)

func GetUserIDFromMetadata(username string) string {
	url := fmt.Sprintf("%s/%s/?__a=1", BaseUrl, username)
	body := gorammaHTTPRequest(url)

	var userMetadata = new(structs.IGAPIUserMetadata)
	err := json.Unmarshal(body, &userMetadata)

	if err != nil {
		log.Fatal(err)
	}

	return userMetadata.GraphQL.User.ID
}

func GetMediaDetailFromShortcode(shortcode string) *structs.IGAPIMediaDetail {
	// Example URL: https://www.instagram.com/p/B-AlSmXAYFM/?__a=1
	url := fmt.Sprintf("%s/p/%s/?__a=1", BaseUrl, shortcode)
	body := gorammaHTTPRequest(url)

	var mediaDetail = new(structs.IGAPIMediaDetail)
	err := json.Unmarshal(body, &mediaDetail)

	if err != nil {
		log.Fatal(err)
	}

	return mediaDetail
}

func GetUserTimelineMedia(userID string, endCursor string) *structs.IGAPITimeline {
	url := buildGorammaNextPageURL(userID, endCursor)
	body := gorammaHTTPRequest(url)

	var timeline = new(structs.IGAPITimeline)
	err := json.Unmarshal(body, &timeline)

	if err != nil {
		log.Fatal(err)
	}

	return timeline
}
