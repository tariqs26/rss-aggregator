package scraper

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urlToFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: time.Second * 10,
	}

	response, err := httpClient.Get(url)

	rssFeed := RSSFeed{}

	if err != nil {
		return rssFeed, err
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)

	if err != nil {
		return rssFeed, err
	}

	err = xml.Unmarshal(data, &rssFeed)

	if err != nil {
		return rssFeed, err
	}

	return rssFeed, nil
}
