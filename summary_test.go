package main

import (
	"os"
	"testing"

	"github.com/mattermost/mattermost-server/model"
)

var (
	URL         string // your mattermost url
	CHANNNEL_ID string // target channel id
	USERNAME    string // your user name
	PASSWORD    string // your password
)

func setEnv() {
	URL = os.Getenv("MATTERMOST_SERVER_URL")
	CHANNNEL_ID = os.Getenv("CHANNNEL_ID")
	USERNAME = os.Getenv("MATTERMOST_USERNAME")
	PASSWORD = os.Getenv("MATTERMOST_PASSWD")
}

// Can create the mattermost client.
func TestCreateClient(t *testing.T) {
	setEnv()
	_, err := createClient(URL, USERNAME, PASSWORD)
	if err != nil {
		t.Fatal(err)
	}
}

// Can post.
func TestCreatePost(t *testing.T) {
	setEnv()
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

func TestGetPosts(t *testing.T) {
	setEnv()
	client, _ := createClient(URL, USERNAME, PASSWORD)
	pl, err := getPosts(client, CHANNNEL_ID)
	if err != nil {
		t.Fatal(err)
	}
	posts := PostToSlice(pl)
	for _, post := range posts {
		t.Log(post.Message)
	}
}

func TestBuildPostMessage(t *testing.T) {
	s := map[string]int{}
	s["userA"] = 10
	s["userB"] = 9
	s["userC"] = 8
	s["userD"] = 7
	s["userE"] = 6
	s["hogehoge"] = 5
	s["fugafuga"] = 0
	t.Log(buildPostMessage(s))
}

func TestSummary(t *testing.T) {
	req := &Request{}
	req.ChannelID = CHANNNEL_ID
	res := Summary(*req)
	t.Log(res.Text)
}

// Util
func PrintError(err *model.AppError) {
	println("\tError Details:")
	println("\t\t" + err.Message)
	println("\t\t" + err.Id)
	println("\t\t" + err.DetailedError)
}
