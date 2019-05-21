package main

import (
	"fmt"
	"log"
	"sort"

	mm "github.com/mattermost/mattermost-server/model"
)

var client *mm.Client4

// Max API fetch count.
const FETCHCOUNT = 100

type summary map[string]int

// Summary is call from plug-in entry point.
func Summary() *Response {
	res := &Response{}
	client, err := createClient("", "", "")
	if err != nil {
		log.Fatal(err)
	}
	getPosts(client, "")
	return res
}

// create WebAPI Client
func createClient(url string, userName string, passwd string) (*mm.Client4, error) {
	client = mm.NewAPIv4Client(url)
	_, res := client.Login(userName, passwd)
	if res.Error != nil {
		return nil, createFmtError(res)
	}
	return client, nil
}

// getPosts
func getPosts(c *mm.Client4, channel string) (*mm.PostList, error) {
	post, res := c.GetPostsForChannel(channel, 0, FETCHCOUNT, "")
	if res.Error != nil {
		return nil, createFmtError(res)
	}

	return post, nil
}

// util.
func createFmtError(e *mm.Response) error {
	return fmt.Errorf("Error :%s, [code: %d]", e.Error.Message, e.StatusCode)
}

// PostToSlice --> Not required for the latest version.
func PostToSlice(pl *mm.PostList) []*mm.Post {
	var posts []*mm.Post
	for _, id := range pl.Order {
		posts = append(posts, pl.Posts[id])
	}
	return posts
}

// Summarize action.
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

var (
	icons = map[int]string{
		0: ":confetti_ball:",
		1: ":tada:",
		2: ":100:",
		3: ":ideograph_advantage:",
		4: ":beginner:",
	}
)

// Build post message from summary data.
func buildPostMessage(m summary) string {
	var header = "@here Weekly awards! \n | order | user | icon |\n| :-: | :-- | :-: |\n"

	values := []int{}
	var rank string
	for _, v := range m {
		values = append(values, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	for i := 0; i < 5; i++ {
		if len(values) <= 0 {
			break
		}
		arr := m.valueToKey(values[len(values)-1])
		for _, v := range arr {
			rank += fmt.Sprintf("| %d | %s | %s |\n", i+1, v, icons[i])
		}
		values = popSlice(values)
	}
	return fmt.Sprintf("%s%s", header, rank)
}

// Find key from value.
func (s summary) valueToKey(idx int) []string {
	result := []string{}
	for i, v := range s {
		if v == idx {
			result = append(result, i)
		}
	}
	return result
}

func popSlice(slice []int) []int {
	slice = slice[:len(slice)-1]
	return slice
}
