package model

import (
	"testing"
	"thanks/model"
)

func TestParse(t *testing.T) {
	const param = "!channel_id=3zkcxrd313gfdqipfx41nj5sny&channel_name=test_ishibashifutos&command=%2Fthx&response_url=http%3A%2F%2Flocalhost%3A8080%2Fhooks%2Fcommands%2Faq9m7jg4qjfhbpxmx65tez5rmy&team_domain=test&team_id=t6rwj1ncsjn9xcnm3n5ynrgbyc&text=%40bash+%22thank+you%21%22&token=tc1i8gzp43bhbcss1dqedu47he&trigger_id=eG9wZTM5bTNzZmZvN3llaDN5czF6ejVmb3I6a3M1MTd5NmFvN2ZhamNzNHlodWZkMzZuM3k6MTU1ODMxMzkyODk5MzpNRVVDSVFEaTZDWWdlWmJMZytuNXVRblVLbHhxUVJWdzBiK1RWblg0Y3NES0tEUUNkUUlnTUZqTFVaMzJ3bDF6UHY4NmJVK25WNlFDS08vd3VkRFNDUEVKM2JwNHhKaz0%3D&user_id=ks517y6ao7fajcs4yhufd36n3y&user_name=bash"
	r, err := model.ParseRequest(param)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
