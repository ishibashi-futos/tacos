package main

import (
	"fmt"
	"log"

	mm "github.com/mattermost/mattermost-server/model"

	"thanks/model"
	"thanks/util"
)

func ephemeralPost(destination string) model.Response {
	model := new(model.Response)
	model.ResponseType = "ephemeral"
	model.UserName = "お令和"
	model.Text = fmt.Sprintf("了解トク! %s さんに感謝を伝えたトク！", destination)
	return *model
}

// Max API fetch count.
const URL = "@YOURENV"
const UN = "@YOURENV"
const PW = "@YOURENV"

func thanksMessagePost(destination string, message string, channelID string) error {
	client, err := util.CreateClient(URL, UN, PW)
	if err != nil {
		return err
	}
	defer func() {
		if _, res := client.Logout(); res.Error != nil {
			log.Fatal(util.CreateFmtError(res))
		}
	}()
	newPost := mm.Post{
		UserId:    "tokumaro",
		ChannelId: channelID,
		Message:   fmt.Sprintf("# :o_reiwa: \n %s %s", destination, message),
	}
	post, res := client.CreatePost(&newPost)
	if res.Error != nil {
		return util.CreateFmtError(res)
	}
	log.Println(post.ToJson())
	return nil
}
