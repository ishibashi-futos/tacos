package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mattermost/mattermost-server/plugin"
)

type HelloWorldPlugin struct {
	plugin.MattermostPlugin
}

type Response struct {
	ResponseType string `json:"response_type"`
	UserName     string `json:"tacos"`
	Text         string `json:"text"`
	IconUrl      string `json:"icon_url"`
}

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

func (p *HelloWorldPlugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	bytes, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var req Request
	err := json.Unmarshal(bytes, &req)

	model := new(Response)
	model.ResponseType = "in_channel"
	if err != nil {
		model.Text = "@here thank you !" + err.Error()
	} else {
		model.Text = "@here thank you !" + req.Text
	}
	model.UserName = "tacos"
	model.IconUrl = "https://www.mattermost.org/wp-content/uploads/2016/04/icon.png"
	json, _ := json.Marshal(&model)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

// This example demonstrates a plugin that handles HTTP requests which respond by greeting the
// world.
func main() {
	plugin.ClientMain(&HelloWorldPlugin{})
}
