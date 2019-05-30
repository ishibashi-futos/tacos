package main

import (
	"fmt"

	mm "github.com/mattermost/mattermost-server/model"
)

// Max API fetch count.
const FETCHCOUNT = 100
const MM_URL = "@YOURENV"
const UN = "@YOURENV"
const PW = "@YOURENV"

// create WebAPI Client
func createClient(url string, userName string, passwd string) (*mm.Client4, error) {
	client = mm.NewAPIv4Client(url)
	_, res := client.Login(userName, passwd)
	if res.Error != nil {
		return nil, createFmtError(res)
	}
	return client, nil
}

// util.
func createFmtError(e *mm.Response) error {
	return fmt.Errorf("Error :%s, [code: %d]", e.Error.Message, e.StatusCode)
}
