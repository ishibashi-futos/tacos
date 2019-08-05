package model

import (
	"errors"
	"strings"
)

type Request struct {
	ChannelID   string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	Command     string `json:"command"`
	ResponseURL string `json:"response_url"`
	TeamDomain  string `json:"team_domain"`
	TeamID      string `json:"team_id"`
	Text        string `json:"text"`
	Token       string `json:"token"`
	UserID      string `json:"user_id"`
	UserName    string `json:"user_name"`
}

func (r Request) TextToMessage() (s0 string, s1 string, err error) {
	idx := strings.Index(r.Text, " ")
	if idx == -1 {
		err = errors.New("")
	}
	s0 = r.Text[:idx]
	s1 = r.Text[idx+1:]

	return
}
