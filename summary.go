package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	mm "github.com/mattermost/mattermost-server/model"
)

var client *mm.Client4

type summary map[string]int

// Summary is call from plug-in entry point.
func Summary(req Request) *Response {
	res := &Response{}
	client, err := createClient(MM_URL, UN, PW)
	if err != nil {
		log.Fatal(err)
	}
	pl, _ := getPosts(client, req.ChannelID)
	posts := PostToSlice(pl)
	s := summarize(posts)
	if len(s) != 0 {
		res.Text = buildPostMessage(s)
	}
	return res
}


// getPosts
func getPosts(c *mm.Client4, channel string) (*mm.PostList, error) {
	post, res := c.GetPostsForChannel(channel, 0, FETCHCOUNT, "")
	if res.Error != nil {
		return nil, createFmtError(res)
	}

	return post, nil
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
func summarize(posts []*mm.Post) summary {
	var s map[string]int = map[string]int{}
	for _, post := range posts {
		user := findUserFromPost(post.Message)
		if user == "" {
			break
		}
		if _, ok := s[user]; ok {
			s[user]++
		} else {
			s[user] = 0
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
	var header = "@here Weekly awards! \n | order | user | icon |\n| :-: | :-- | :-: |\n\n"

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

// find target user
func findUserFromPost(message string) string {
	idx := strings.Index(message, " ")
	if idx == -1 {
		return ""
	}
	return message[:idx]
}
