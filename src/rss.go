package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title     string    `xml:"title"`
		Link      string    `xml:"link"`
		Language  string    `xml:"language"`
		Copyright string    `xml:"copyright"`
		Item      []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

// `urlToFeed` takes a URL as input and returns an RSSFeed struct and an error
// It makes an HTTP GET request to the specified URL, reads the response body,
// and unmarshals the XML data into the RSSFeed struct
func urlToFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	// Send an HTTP GET request to the specified URL using the client
	response, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}

	defer response.Body.Close()

	// Read the response body into a byte slice
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return RSSFeed{}, err
	}

	rssFeed := RSSFeed{}
	// Unmarshal the XML data from the byte slice into the RSSFeed struct
	err = xml.Unmarshal(bytes, &rssFeed)
	if err != nil {
		return RSSFeed{}, err
	}
	return rssFeed, nil
}
