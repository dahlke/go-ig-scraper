package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const BASE_URL = "https://www.instagram.com"

type IGUserTimelineMediaEdgeNodeLocation struct {
	Name string `json:"name"`
}

type IGUserTimelineMediaEdgeNode struct {
	DisplayURL string                              `json:"display_url"`
	Timestamp  string                              `json:"taken_at_timestamp"`
	Location   IGUserTimelineMediaEdgeNodeLocation `json:"location"`
}

type IGUserTimelineMediaEdges struct {
	Node IGUserTimelineMediaEdgeNode `json:"node"`
}

type IGUserTimelineMediaPageInfo struct {
	HasNextPage bool   `json:"has_next_page"`
	EndCursor   string `json:"end_cursor"`
}

type IGUserTimelineMedia struct {
	Edges    []IGUserTimelineMediaEdges  `json:"edges"`
	PageInfo IGUserTimelineMediaPageInfo `json:"page_info"`
}

type IGUser struct {
	ID            string              `json:"id"`
	TimelineMedia IGUserTimelineMedia `json:"edge_owner_to_timeline_media"`
}

type IGUserGraphQL struct {
	User IGUser `json:"user"`
}

type IGUserMetadata struct {
	GraphQL IGUserGraphQL `json:"graphql"`
}

func getUserMetadataFromInstagramURL(url string) *IGUserMetadata {
	httpClient := http.Client{Timeout: time.Second * 2}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "los-jabronis")

	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var userMetadata = new(IGUserMetadata)
	err = json.Unmarshal(body, &userMetadata)

	return userMetadata
}

func main() {
	// I get by, with a little help, from this blog:
	// http://go-colly.org/articles/how_to_scrape_instagram/
	username := "eklhad"
	userInfoURL := fmt.Sprintf("%s/%s/?__a=1", BASE_URL, username)
	userMetadata := getUserMetadataFromInstagramURL(userInfoURL)

	if userMetadata.GraphQL.User.TimelineMedia.PageInfo.HasNextPage {
		endCursor := userMetadata.GraphQL.User.TimelineMedia.PageInfo.EndCursor
		userID := userMetadata.GraphQL.User.ID
		queryMediaVars := url.QueryEscape(fmt.Sprintf("{\"id\":\"%s\",\"first\":50,\"after\":\"%s\"}", userID, endCursor))
		nextPageURL := fmt.Sprintf("%s/%s&variables=%s", BASE_URL, "graphql/query/?query_id=17888483320059182", queryMediaVars)
		// userMetadata = getUserMetadataFromInstagramURL(nextPageURL)
		fmt.Println("QUERY MEDIA URL", queryMediaURL)
	}
}
