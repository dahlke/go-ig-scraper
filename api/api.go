package api

import (
	"encoding/json"
	"fmt"

	structs "github.com/dahlke/goramma/structs"
	log "github.com/sirupsen/logrus"
)

func convertMediaDetail(mediaTimelineSlice []structs.IGAPIMediaDetail) []structs.InstagramMedia {
	allConvertedMedia := []structs.InstagramMedia{}

	for _, m := range mediaTimelineSlice {
		convertedMedia := structs.InstagramMedia{
			ShortCode: m.GraphQL.Media.Shortcode,
			Timestamp: m.GraphQL.Media.Timestamp,
			Location:  m.GraphQL.Media.Location.Name,
			URL:       m.GraphQL.Media.DisplayURL,
		}
		allConvertedMedia = append(allConvertedMedia, convertedMedia)
	}

	return allConvertedMedia
}

func getMediaDetailFromShortcode(shortcode string) *structs.IGAPIMediaDetail {
	// Example URL: https://www.instagram.com/p/B-AlSmXAYFM/?__a=1
	url := fmt.Sprintf("%s/p/%s/?__a=1", BaseURL, shortcode)
	body := gorammaHTTPRequest(url)

	var mediaDetail = new(structs.IGAPIMediaDetail)
	err := json.Unmarshal(body, &mediaDetail)

	if err != nil {
		log.Fatal(err)
	}

	return mediaDetail
}

// GetUserIDFromMetadata retrieves a User ID from the supplied username.
func GetUserIDFromMetadata(username string) string {
	url := fmt.Sprintf("%s/%s/?__a=1", BaseURL, username)
	body := gorammaHTTPRequest(url)

	var userMetadata = new(structs.IGAPIUserMetadata)
	err := json.Unmarshal(body, &userMetadata)

	if err != nil {
		log.Fatal(err)
	}

	return userMetadata.GraphQL.User.ID
}

// GetUserTimelineMedia retrieves all data starting after the endCursor supplied.
func GetUserTimelineMedia(userID string, endCursor string) ([]structs.InstagramMedia, bool, string) {
	url := buildGorammaNextPageURL(userID, endCursor)
	body := gorammaHTTPRequest(url)

	var timeline = new(structs.IGAPITimeline)
	err := json.Unmarshal(body, &timeline)

	if err != nil {
		log.Fatal(err)
	}

	timelineEdges := timeline.Data.User.Media.Edges
	hasNextPage := timeline.Data.User.Media.PageInfo.HasNextPage
	endCursor = timeline.Data.User.Media.PageInfo.EndCursor

	var mediaTimelineSlice []structs.IGAPIMediaDetail

	for i, edge := range timelineEdges {
		shortcode := edge.Node.Shortcode

		log.Info(fmt.Sprintf("Getting the details for media %d of %d...", i+1, len(timelineEdges)))
		mediaDetail := getMediaDetailFromShortcode(shortcode)
		mediaTimelineSlice = append(mediaTimelineSlice, *mediaDetail)
	}

	convertedMedia := convertMediaDetail(mediaTimelineSlice)
	return convertedMedia, hasNextPage, endCursor
}
