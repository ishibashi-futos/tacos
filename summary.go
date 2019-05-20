package main

import (
	"fmt"
	"log"
	"sort"

	mm "github.com/mattermost/mattermost-server/model"
)

var client *mm.Client4
type summary map[string]int

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

func summarize(posts []mm.Post) summary {
	var s map[string]int
	for _, post := range posts {
		if _, ok := s[post.Message]; ok {
			s[post.Message]++
		} else {
			s[post.Message] = 0
		}
	}
	return s
}

func buildPostMessage(m summary) string {

	var header = "| order | user | icon |\n| :-: | :-- | :-: |"

	var values = []int{}
	for _, v := range m {
		values = append(values, v)
	}
	sort.Sort(sort.IntSlice(values))
	return fmt.Sprintf("%s, %v", header, values)
}