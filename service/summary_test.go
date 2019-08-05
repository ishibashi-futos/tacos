package service

import (
	"testing"

	"thanks/model"
	"thanks/repository"
	"thanks/service"
)

func TestCreatePost(t *testing.T) {
	req := &model.Request{}
	res := service.Summary(req)
	t.Log(res.Text)
}

func TestBuildPostMessage(t *testing.T) {
	s := make(map[string]int)
	s["userA"] = 10
	s["userB"] = 9
	s["userC"] = 8
	s["userD"] = 7
	s["userE"] = 6
	s["hogehoge"] = 5
	s["fugafuga"] = 0
	t.Log(service.BuildPostMessage(s, repository.NewPeriod("Weekly").String()))
}
