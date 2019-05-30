package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/mattermost/mattermost-server/plugin"
)

type TacosPlugin struct {
	plugin.MattermostPlugin
}

func (p *TacosPlugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	bytes, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	req, _ := parseRequest(string(bytes))

	s0, s1, _ := req.textToMessage()
	if err := thanksMessagePost(s0, s1, req.ChannelID); err != nil {
		log.Fatal(err)
	}
	model := ephemeralPost(s0)

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
