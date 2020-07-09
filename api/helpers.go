package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	log "github.com/sirupsen/logrus"
)

func gorammaHTTPRequest(igURL string) []byte {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

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

func buildGorammaNextPageURL(userID string, endCursor string) string {
	// Example URL:
	// https://www.instagram.com/graphql/query/?query_id=17888483320059182&variables=%7B%22id%22%3A%2211321561%22%2C%22first%22%3A12%2C%22after%22%3A%22%22%7D

	queryMediaVars := url.QueryEscape(fmt.Sprintf("{\"id\":\"%s\",\"first\":50,\"after\":\"%s\"}", userID, endCursor))
	nextPageURL := fmt.Sprintf("%s/%s&variables=%s", BaseURL, "graphql/query/?query_id=17888483320059182", queryMediaVars)

	return nextPageURL
}
