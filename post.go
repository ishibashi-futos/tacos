package main

import (
	"fmt"
	"log"

	mm "github.com/mattermost/mattermost-server/model"
)

func ephemeralPost(destination string) Response {
	model := new(Response)
	model.ResponseType = "ephemeral"
	model.UserName = "お令和"
	model.Text = fmt.Sprintf("了解トク! %s さんに感謝を伝えたトク！", destination)
	return *model
}

func thanksMessagePost(destination string, message string, channelID string) error {
	client, err := createClient("XXX", "XXX", "XXX")
	if err != nil {
		return err
	}
	defer func() {
		if _, res := client.Logout(); res.Error != nil {
			log.Fatal(createFmtError(res))
		}
	}()
	newPost := mm.Post{
		UserId:    "tokumaro",
		ChannelId: channelID,
		Message:   fmt.Sprintf("# :o_reiwa: \n %s %s", destination, message),
	}
	post, res := client.CreatePost(&newPost)
	if res.Error != nil {
		return createFmtError(res)
	}
	log.Println(post.ToJson())
	return nil
}
