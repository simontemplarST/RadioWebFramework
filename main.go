package main

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/bennof/accessDBwE"
	"github.com/gohugoio/hugo/parser"
	"github.com/gohugoio/hugo/parser/metadecoders"
	_ "github.com/mattn/go-sqlite3"
)

type Configuration struct {
	DBType          string `json:"dbType"`
	Database        string `json:"database"`
	Table           string `json:"table"`
	ContentColumn   string `json:"contentColumn"`
	TitleColumn     string `json:"titleColumn"`
	Description     string `json:"description"`
	Link            string `json:"link"`
	Author          string `json:"author"`
	Email           string `json:"email"`
	RSSFileName     string `json:"rssFileName"`
	MarkdownFolder  string `json:"markdownFolder"`
	StatusRSSFile   string `json:"statusRSSFile"`
	OnAirText       string `json:"onAirText"`
	OffAirText      string `json:"offAirText"`
	IntervalMinutes int    `json:"intervalMinutes"`
	HugoPostTitle   string `json:"hugoPostTitle"`
}

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Author      string   `xml:"author"`
	Category    string   `xml:"category"`
	PubDate     string   `xml:"pubDate"`
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Items       []Item   `xml:"item"`
}

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

func main() {
	// Load Configuration
	config := loadConfig("config.json")
	// Open DB connection
	var db *sql.DB
	var err error
	if config.DBType == "access" {
		db, err = accessDBwE.Open("adodb", "Provider=Microsoft.ACE.OLEDB.12.0;Data Source="+config.Database)
	} else if config.DBType == "sqlite" {
		db, err = sql.Open("sqlite3", config.Database)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Get Data from DB
	rows, err := db.Query("SELECT * FROM " + config.Table)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Loop over DB rows, create Items and write to Hugo
	var items []Item
	for rows.Next() {
		// Extract data from row
		var id int
		var content string
		err = rows.Scan(&id, &content)
		if err != nil {
			log.Fatal(err)
		}

		// Create an Item and append to items
		item := Item{
			Title:       fmt.Sprintf("%s %d", config.TitleColumn, id),
			Link:        config.Link,
			Description: content,
			Author:      config.Author,
			PubDate:     time.Now().Format(time.RFC822),
		}
		items = append(items, item)

		// Create a Hugo post
		writeToHugo(item, config.MarkdownFolder, config.HugoPostTitle)

		// Wait for interval
		time.Sleep(time.Duration(config.IntervalMinutes) * time.Minute)

		// Check if next episode is ready
		status := checkStatus(db, config.Table, config.ContentColumn)
		writeStatusRSS(status, config)

		// If on air, publish next episode
		if status == config.OnAirText {
			publishNextEpisode(db, config.Table, config.ContentColumn)
		}
	}

	// Write Items to RSS
	rss := &RSS{
		Version: "2.0",
		Channel: Channel{
			Title:       config.TitleColumn,
			Link:        config.Link,
			Description: config.Description,
			Items:       items,
		},
	}
	rssData, err := xml.MarshalIndent(rss, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// Write RSS data to file
	err = ioutil.WriteFile(config.RSSFileName, rssData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func writeToHugo(item Item, folder string, title string) {
	// Create a new File
	newpath := filepath.Join(folder, fmt.Sprintf("%s.md", item.Title))
	newFile, err := os.Create(newpath)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// Create a new Page
	page, err := parser.NewPageParser().Parse(strings.NewReader(""))
	if err != nil {
		log.Fatal(err)
	}

	// Set page metadata
	page.FrontMatter["title"] = title
	page.FrontMatter["date"] = time.Now().Format(time.RFC3339)
	page.FrontMatter["draft"] = false

	// Set page content
	page.Content = item.Description

	// Write to File
	err = page.Render(newFile, metadecoders.FormatJSON)
	if err != nil {
		log.Fatal(err)
	}
}

func checkStatus(db *sql.DB, table string, contentColumn string) string {
	rows, err := db.Query("SELECT " + contentColumn + " FROM " + table + " LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		return "On"
	} else {
		return "Off"
	}
}

func writeStatusRSS(status string, config Configuration) {
	// Create a new status Item
	item := Item{
		Title:       "Status",
		Link:        config.Link,
		Description: status,
		Author:      config.Author,
		PubDate:     time.Now().Format(time.RFC822),
	}

	// Write status Item to RSS
	rss := &RSS{
		Version: "2.0",
		Channel: Channel{
			Title:       "Status",
			Link:        config.Link,
			Description: "Status RSS",
			Items:       []Item{item},
		},
	}
	rssData, err := xml.MarshalIndent(rss, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// Write RSS data to file
	err = ioutil.WriteFile(config.StatusRSSFile, rssData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func publishNextEpisode(db *sql.DB, table string, contentColumn string) {
	_, err := db.Exec("UPDATE " + table + " SET " + contentColumn + " = 'Off' WHERE " +
		contentColumn + " = 'On' LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
}

func LoadOpen() {
	// Load Configuration
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	// Open Database
	db, err := accessdbwe.Open("adodb", "Provider=Microsoft.ACE.OLEDB.12.0;Data Source="+config.DatabasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create RSS Feed
	createRSSFeed(db, config)
}
