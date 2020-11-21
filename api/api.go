package api

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	structs "github.com/dahlke/goramma/structs"
	log "github.com/sirupsen/logrus"
)

// InstagramTimestampFmt is used as a format for parsing dates from the Instagram Graph API
const InstagramTimestampFmt = "2006-01-02T15:04:05-0700"

// GetUserMetadata will retrieve the user metadata from Instagram Graph API.
func GetUserMetadata(instagramToken string) structs.InstagramUserMetadata {
	url := fmt.Sprintf("%s/me?fields=id,username&access_token=%s", BaseURL, instagramToken)
	body := gorammaHTTPRequest(url)

	var userMetadata = new(structs.InstagramUserMetadata)
	err := json.Unmarshal(body, &userMetadata)
	if err != nil {
		log.Fatal(err)
	}

	return *userMetadata
}

// GetUserMedia will retrieve a page of media results from the Instagram Graph API.
func GetUserMedia(instagramToken string, endCursor string) structs.InstagramUserMedia {
	url := fmt.Sprintf("%s/me/media?fields=id,media_type,media_url,permalink,username,timestamp,caption&limit=50&access_token=%s&after=%s", BaseURL, instagramToken, endCursor)
	body := gorammaHTTPRequest(url)

	var userMedia = new(structs.InstagramUserMedia)
	err := json.Unmarshal(body, &userMedia)
	if err != nil {
		log.Fatal(err)
	}

	for i := range userMedia.Data {
		timestamp, err := time.Parse(InstagramTimestampFmt, userMedia.Data[i].Timestamp)
		if err != nil {
			log.Fatal(err)
		}
		unixTimestampStr := strconv.FormatInt(timestamp.Unix(), 10)
		userMedia.Data[i].Timestamp = unixTimestampStr
	}

	return *userMedia
}
