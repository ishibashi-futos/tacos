package service

import (
	"fmt"
	"sort"

	"thanks/model"
	"thanks/repository"
)

type SummaryData map[string]int

// Summary is call from plug-in entry point.
func Summary(req *model.Request) *model.Response {
	res := &model.Response{}
	repo, _ := repository.NewThanksRepository()
	defer repo.Close()
	var p repository.Period = repository.Week
	data, _ := repo.Summarize(p)
	if len(data) != 0 {
		res.Text = BuildPostMessage(data, p.String())
	}
	return res
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
func BuildPostMessage(m SummaryData, p string) string {
	var header = fmt.Sprintf("@here %s awards! \n | order | user | icon |\n| :-: | :-- | :-: |\n", p)
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
func (s SummaryData) valueToKey(idx int) []string {
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
