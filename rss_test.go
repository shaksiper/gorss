package main

import (
	"reflect"
	"testing"
)

func TestEmptyURL(t *testing.T) {
	URL := ""
	rssFeed, err := urlToFeed(URL)
	if err == nil {
		t.Fatal("Empty URL should return an error")
	}
	if !reflect.ValueOf(rssFeed).IsZero() {
		t.Fatal("RSSFeed should be empty to indicate error")
	}
}

func TestNonRSSURL(t *testing.T) {
	URL := "http://localhost"
	rssFeed, err := urlToFeed(URL)
	if err == nil {
		t.Fatal("Non RSS URL should return an error")
	}
	if !reflect.ValueOf(rssFeed).IsZero() {
		t.Fatal("RSSFeed should be empty to indicate error")
	}
}

func TestCorrectRSSURL(t *testing.T) {
    URL := "https://feeds.megaphone.fm/newheights"
	rssFeed, err := urlToFeed(URL)
	if err != nil {
		t.Fatal("Valid RSS URL should not return error")
	}
	if reflect.ValueOf(rssFeed).IsZero() {
		t.Fatal("RSSFeed struct returned empty from valid URL")
	}

	// t.Log(rssFeed)
}
