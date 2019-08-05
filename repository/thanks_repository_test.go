package repository

import (
	"os"
	"testing"

	"thanks/repository"
)

var repo *repository.ThanksRepository

func TestMain(m *testing.M) {
	repo, _ = repository.NewThanksRepository()
	defer repo.Close()
	if ret := m.Run(); ret != 0 {
		os.Exit(ret)
	}
}

func TestSave(t *testing.T) {
	repo.Save("tokumaro", "test", "Hello, World!")
}

func TestSummarize(t *testing.T) {
	var data map[string]int
	var p repository.Period = repository.Week
	data, err := repo.Summarize(p)
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range data {
		t.Log(k, v)
	}
}