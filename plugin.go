package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mattermost/mattermost-server/plugin"
)

type TacosPlugin struct {
	plugin.MattermostPlugin
}

type Response struct {
	ResponseType string `json:"response_type"`
	UserName     string `json:"tacos"`
	Text         string `json:"text"`
	IconURL      string `json:"icon_url"`
}

func (p *TacosPlugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	bytes, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	req, _ := parseRequest(string(bytes))

	model := new(Response)
	model.ResponseType = "in_channel"
	s0, s1, _ := req.textToMessage()
	model.Text = fmt.Sprintf("%s %s", s0, s1)
	model.UserName = "tacos"
	model.IconURL = "https://www.mattermost.org/wp-content/uploads/2016/04/icon.png"
	json, _ := json.Marshal(&model)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func (r Request) textToMessage() (s0 string, s1 string, err error) {
	idx := strings.Index(r.Text, " ")
	if idx == -1 {
		err = errors.New("")
	}
	s0 = r.Text[:idx]
	s1 = r.Text[idx+1:]

	return
}

func main() {
	plugin.ClientMain(&TacosPlugin{})
}
