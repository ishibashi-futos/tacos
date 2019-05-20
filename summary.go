package main

import (
	"fmt"

	mm "github.com/mattermost/mattermost-server/model"
)

func summary() *Response {
	res := &Response{}

	return res
}

func createClient(url string, userName string, passwd string) (*mm.Client4, error) {
	client := mm.NewAPIv4Client(url)
	_, res := client.Login(userName, passwd)
	if res.Error != nil {
		return nil, fmt.Errorf("Error :%s, [code: %d], [detail %s]", res.Error.Message, res.StatusCode, res.Error.DetailedError)
	}
	return client, nil
}
