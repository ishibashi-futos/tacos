package main

import (
	"fmt"
	"log"

	mm "github.com/mattermost/mattermost-server/model"
)

var client *mm.Client4

func Summary() *Response {
	res := &Response{}
	client, err := createClient("", "", "")
	if err != nil {
		log.Fatal(err)
	}
	getPosts(client, "")
	return res
}

func createClient(url string, userName string, passwd string) (*mm.Client4, error) {
	client = mm.NewAPIv4Client(url)
	_, res := client.Login(userName, passwd)
	if res.Error != nil {
		return nil, createFmtError(res)
	}
	return client, nil
}

func getPosts(c *mm.Client4, channel string) (*mm.PostList, error){
	post, res := c.GetPostsForChannel(channel, 0, 10, "")
	if res.Error != nil {
		return nil, createFmtError(res)
	}

	return post, nil
}

func createFmtError(e *mm.Response) error {
	return fmt.Errorf("Error :%s, [code: %d]", e.Error.Message, e.StatusCode)
}

func PostToSlice(pl *mm.PostList) []*mm.Post {
	var posts []*mm.Post
	for _, id := range pl.Order {
		posts = append(posts, pl.Posts[id])
	}
	return posts
}
