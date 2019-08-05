package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mattermost/mattermost-server/plugin"

	"thanks/model"
	"thanks/repository"
)

type ThanksPlugin struct {
	plugin.MattermostPlugin
}

func (p *ThanksPlugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	bytes, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	req, _ := model.ParseRequest(string(bytes))

	s0, s1, _ := req.TextToMessage()
	repo, err := repository.NewThanksRepository()
	if err != nil {
		log.Fatal(err)
	}
	defer repo.Close()
	repo.Save(req.UserID, s0, s1)
	if err := thanksMessagePost(s0, s1, req.ChannelID); err != nil {
		log.Fatal(err)
	}
	model := ephemeralPost(s0)

	json, _ := json.Marshal(&model)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func main() {
	plugin.ClientMain(&ThanksPlugin{})
}
