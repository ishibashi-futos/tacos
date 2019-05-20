package main

import (
	"testing"

	"github.com/mattermost/mattermost-server/model"
)

const (
	URL         = "<@your mattermost url>" // your mattermost url
	CHANNNEL_ID = "<@target channel id>"   // target channel id
	USERNAME    = "<@your user name>"      // your user name
	PASSWORD    = "<@your password>"       // your password
)

/*
Can create the mattermost client.
*/
func TestCreateClient(t *testing.T) {

	_, err := createClient(URL, USERNAME, PASSWORD)
	if err != nil {
		t.Fatal(err)
	}
}

/*
Can post.
*/
func TestCreatePost(t *testing.T) {
	// reference
	// https://github.com/mattermost/mattermost-bot-sample-golang/blob/master/bot_sample.go
	client, err := createClient(URL, USERNAME, PASSWORD)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Logout()

	post := &model.Post{}
	post.ChannelId = CHANNNEL_ID
	post.Message = "Hello, World!"
	if _, res := client.CreatePost(post); res.Error != nil {
		t.Fatal("We failed to send a message to the logging channel")
		PrintError(res.Error)
	}
}

func PrintError(err *model.AppError) {
	println("\tError Details:")
	println("\t\t" + err.Message)
	println("\t\t" + err.Id)
	println("\t\t" + err.DetailedError)
}
