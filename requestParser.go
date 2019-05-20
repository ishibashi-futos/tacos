package main

import (
	"net/url"
	"strings"
)

type Request struct {
	ChannelID   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	Command     string `json:"command"`
	ResponseUrl string `json:"response_url"`
	TeamDomain  string `json:"team_domain"`
	TeamId      string `json:"team_id"`
	Text        string `json:"text"`
	Token       string `json:"token"`
	UserID      string `json:"user_id"`
}

func (r Request) String() string {
	return r.Text
}

func parseRequest(text string) (*Request, error) {
	r := &Request{}

	params := strings.Split(strings.Replace(text, "!", "", 1), "&")

	for _, v := range params {
		kv := strings.Split(v, "=")
		key := kv[0]
		value, _ := url.QueryUnescape(kv[1])
		switch key {
		case "channel_id":
			r.ChannelID = value
		case "channel_name":
			r.ChannelName = value
		case "command":
			r.Command = value
		case "text":
			r.Text = value
		}
	}

	return r, nil
}
