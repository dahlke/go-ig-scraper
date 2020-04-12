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

/* IG User Timeline Structs */
type IGTimeline struct {
	Data IGTimelineData `json:"data"`
}

type IGTimelineData struct {
	User IGTimelineUser `json:"user"`
}

type IGTimelineUser struct {
	User IGTimelineMedia `json:"edge_owner_to_timeline_media"`
}

type IGTimelineMedia struct {
	Count    int                     `json:"count"`
	Edges    []IGTimelineMediaEdge   `json:"edges"`
	PageInfo IGTimelineMediaPageInfo `json:"page_info"`
}

type IGTimelineMediaPageInfo struct {
	HasNextPage bool   `json:"has_next_page"`
	EndCursor   string `json:"end_cursor"`
}

type IGTimelineMediaEdge struct {
	Node IGTimelineMediaEdgeNode `json:"node"`
}

type IGTimelineMediaEdgeNode struct {
	DisplayURL string                          `json:"display_url"`
	Timestamp  int                             `json:"taken_at_timestamp"`
	Location   IGTimelineMediaEdgeNodeLocation `json:"location"`
}

type IGTimelineMediaEdgeNodeLocation struct {
	Name string `json:"name"`
}

/* IG User Metadata Structs */

type IGUserMetadata struct {
	GraphQL IGUserGraphQL `json:"graphql"`
}

type IGUserGraphQL struct {
	User IGUser `json:"user"`
}

type IGUser struct {
	ID string `json:"id"`
}

func igHTTPRequest(igURL string) []byte {
	httpClient := http.Client{Timeout: time.Second * 2}

	req, err := http.NewRequest(http.MethodGet, igURL, nil)
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

	return body
}

func getUserMetadataFromInstagramURL(url string) *IGUserMetadata {
	body := igHTTPRequest(url)

	var userMetadata = new(IGUserMetadata)
	err := json.Unmarshal(body, &userMetadata)

	if err != nil {
		log.Fatal(err)
	}

	return userMetadata
}

func getUserTimelineMediaFromURL(url string) *IGTimeline {
	body := igHTTPRequest(url)

	var timeline = new(IGTimeline)
	err := json.Unmarshal(body, &timeline)

	if err != nil {
		log.Fatal(err)
	}

	return timeline
}

func main() {
	// I get by, with a little help, from this blog:
	// http://go-colly.org/articles/how_to_scrape_instagram/
	username := "eklhad"
	userInfoURL := fmt.Sprintf("%s/%s/?__a=1", BASE_URL, username)
	userMetadata := getUserMetadataFromInstagramURL(userInfoURL)
	userID := userMetadata.GraphQL.User.ID

	fmt.Println(userID)

	//  after getting the profile metadata, switch to using these urls for retrieving photos
	queryMediaVars := url.QueryEscape(fmt.Sprintf("{\"id\":\"%s\",\"first\":12,\"after\":\"%s\"}", userID, ""))
	nextPageURL := fmt.Sprintf("%s/%s&variables=%s", BASE_URL, "graphql/query/?query_id=17888483320059182", queryMediaVars)
	fmt.Println(nextPageURL)
	timeline := getUserTimelineMediaFromURL(nextPageURL)
	fmt.Println(timeline)
	/*
		fmt.Println("HAS NEXT PAGE", hasNextPage)
		fmt.Println("NUM PICS", len(userMetadata.GraphQL.User.TimelineMedia.Edges))
		fmt.Println("NEXT PAGE URL", nextPageURL)
		hasNextPage := userMetadata.GraphQL.User.TimelineMedia.PageInfo.HasNextPage

		if hasNextPage {
			// TODO: this new url returns a different format.
			endCursor := userMetadata.GraphQL.User.TimelineMedia.PageInfo.EndCursor
			userID := userMetadata.GraphQL.User.ID
			queryMediaVars := url.QueryEscape(fmt.Sprintf("{\"id\":\"%s\",\"first\":12,\"after\":\"%s\"}", userID, endCursor))
			nextPageURL := fmt.Sprintf("%s/%s&variables=%s", BASE_URL, "graphql/query/?query_id=17888483320059182", queryMediaVars)
			userMetadata = getUserMetadataFromInstagramURL(nextPageURL)
			hasNextPage = userMetadata.GraphQL.User.TimelineMedia.PageInfo.HasNextPage
			fmt.Println("HAS NEXT PAGE", hasNextPage)
			fmt.Println("NUM PICS", len(userMetadata.GraphQL.User.TimelineMedia.Edges))
			fmt.Println(userMetadata.GraphQL.User.TimelineMedia.PageInfo)
			fmt.Println("NEXT PAGE URL", nextPageURL)
		}
	*/
}
