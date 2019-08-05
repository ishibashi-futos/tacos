package model

import (
	"net/url"
	"strings"
)

func ParseRequest(text string) (*Request, error) {
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
		case "response_url":
			r.ResponseURL = value
		case "team_domain":
			r.TeamDomain = value
		case "team_id":
			r.TeamID = value
		case "text":
			r.Text = value
		case "token":
			r.Token = value
		case "user_id":
			r.UserID = value
		case "user_name":
			r.UserName = value
		}
	}

	return r, nil
}
