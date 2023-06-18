package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gorilla/feeds"
	_ "github.com/mattn/go-sqlite3"
)

type Post struct {
	Title   string
	Author  string
	Content string
	Date    time.Time
}

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		fmt.Println("Failed to open database:", err)
		return
	}
	defer db.Close()

	row := db.QueryRow("SELECT title, author, content, date FROM posts ORDER BY date DESC LIMIT 1")

	var post Post
	err = row.Scan(&post.Title, &post.Author, &post.Content, &post.Date)
	if err != nil {
		fmt.Println("Failed to read post:", err)
		return
	}

	// Create a markdown file for the post
	content := fmt.Sprintf(`---
title: "%s"
author: "%s"
date: "%s"
---

%s
`, post.Title, post.Author, post.Date.Format(time.RFC3339), post.Content)

	err = ioutil.WriteFile("post.md", []byte(content), 0644)
	if err != nil {
		fmt.Println("Failed to write post:", err)
		return
	}

	// Create an RSS feed
	feed := &feeds.Feed{
		Title:       "My feed",
		Link:        &feeds.Link{Href: "http://example.com"},
		Description: "This is my personal feed!",
		Author:      &feeds.Author{Name: "John Doe", Email: "johndoe@example.com"},
		Created:     time.Now(),
	}

	feed.Items = []*feeds.Item{
		{
			Title:       post.Title,
			Link:        &feeds.Link{Href: "http://example.com/" + post.Title},
			Description: post.Content,
			Author:      &feeds.Author{Name: post.Author},
			Created:     post.Date,
		},
	}

	rss, err := feed.ToRss()
	if err != nil {
		fmt.Println("Failed to generate RSS:", err)
		return
	}

	err = ioutil.WriteFile("feed.rss", []byte(rss), 0644)
	if err != nil {
		fmt.Println("Failed to write RSS:", err)
		return
	}

	// Create a second RSS feed for status
	statusFeed := &feeds.Feed{
		Title:       "Status feed",
		Link:        &feeds.Link{Href: "http://example.com/status"},
		Description: "This feed shows the status of the database.",
		Author:      &feeds.Author{Name: "Status Bot", Email: "status@example.com"},
		Created:     time.Now(),
	}

	status := "On"
	if time.Since(post.Date) > 30*time.Minute {
		status = "Off"
	}

	statusFeed.Items = []*feeds.Item{
		{
			Title:       "Status",
			Link:        &feeds.Link{Href: "http://example.com/status"},
			Description: status,
			Created:     time.Now(),
		},
	}

	s


	statusRss, err := statusFeed.ToRss()
    if err != nil {
        fmt.Println("Failed to generate status RSS:", err)
        return
    }

    err = ioutil.WriteFile("status.rss", []byte(statusRss), 0644)
    if err != nil {
        fmt.Println("Failed to write status RSS:", err)
        return
    }
}