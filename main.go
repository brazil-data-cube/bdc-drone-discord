package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type EmbedFooter struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}

// EmbedAuthor for Embed Author Structure
type EmbedAuthor struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IconURL string `json:"icon_url"`
}

// EmbedField for Embed Field Structure
type EmbedField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Embed is for Embed Structure
type Embed struct {
	Title       string       `json:"title"`
	Description string       `json:"description"`
	URL         string       `json:"url"`
	Color       int          `json:"color"`
	Content     []string     `json:"content"`
	Footer      EmbedFooter  `json:"footer"`
	Author      EmbedAuthor  `json:"author"`
	Fields      []EmbedField `json:"fields"`
	Timestamp   string       `json:"timestamp"`
}

type Payload struct {
	Wait      bool    `json:"wait"`
	Content   string  `json:"content"`
	Username  string  `json:"username"`
	AvatarURL string  `json:"avatar_url"`
	TTS       bool    `json:"tts"`
	Embeds    []Embed `json:"embeds"`
}

func Color() int {
	if os.Getenv("PLUGIN_COLOR") != "" {
		color := strings.Replace(os.Getenv("PLUGIN_COLOR"), "#", "", -1)
		if s, err := strconv.ParseInt(color, 16, 32); err == nil {
			return int(s)
		}
	}

	switch os.Getenv("DRONE_BUILD_STATUS") {
	case "success":
		// green
		return 0x1ac600
	case "failure", "error", "killed":
		// red
		return 0xff3232
	default:
		// yellow
		return 0xffd930
	}
}

func main() {
	var status string

	if os.Getenv("DRONE_BUILD_STATUS") == "success" {
		status = "succeded"
	} else {
		status = "failed"
	}
	var description string
	switch os.Getenv("DRONE_BUILD_EVENT") {
	case "push":
		description = fmt.Sprintf("**%v** pushed to `%v`.", os.Getenv("DRONE_COMMIT_AUTHOR"), os.Getenv("DRONE_COMMIT_BRANCH"))
	case "pull_request":
		description = fmt.Sprintf("**%v** opened a pull request from `%v` to `%v`", os.Getenv("DRONE_COMMIT_AUTHOR"), os.Getenv("DRONE_SOURCE_BRANCH"), os.Getenv("DRONE_COMMIT_BRANCH"))
	case "tag":
		description = fmt.Sprintf("**%v** created tag `%v`.", os.Getenv("DRONE_COMMIT_AUTHOR"), os.Getenv("DRONE_TAG"))
	}
	title := "Build #**%v** on `%v` has %v."
	embed := Embed{
		Title:       fmt.Sprintf(title, os.Getenv("DRONE_BUILD_NUMBER"), os.Getenv("DRONE_REPO"), status),
		Description: description,
		URL:         os.Getenv("DRONE_BUILD_LINK"),
		Color:       Color(),
		Author: EmbedAuthor{
			Name:    os.Getenv("DRONE_COMMIT_AUTHOR"),
			IconURL: os.Getenv("DRONE_COMMIT_AUTHOR_AVATAR"),
		},
	}
	payload := &Payload{
		Embeds: []Embed{embed},
	}

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(payload)
	req, _ := http.NewRequest("POST", os.Getenv("PLUGIN_WEBHOOK"), payloadBuf)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		log.Print(e)
		os.Exit(1)
	}
	defer res.Body.Close()
}
