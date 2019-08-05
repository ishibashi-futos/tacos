package util

import (
	"fmt"

	mm "github.com/mattermost/mattermost-server/model"
)

// create WebAPI Client
func CreateClient(url string, userName string, passwd string) (*mm.Client4, error) {
	client := mm.NewAPIv4Client(url)
	_, res := client.Login(userName, passwd)
	if res.Error != nil {
		return nil, CreateFmtError(res)
	}
	return client, nil
}

// util.
func CreateFmtError(e *mm.Response) error {
	return fmt.Errorf("Error :%s, [code: %d]", e.Error.Message, e.StatusCode)
}
