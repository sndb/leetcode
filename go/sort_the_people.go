package main

import "sort"

func sortPeople(names []string, heights []int) []string {
	people := make([]struct {
		name   string
		height int
	}, len(names))
	for i := 0; i < len(names); i++ {
		people[i].name = names[i]
		people[i].height = heights[i]
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].height > people[j].height
	})
	r := []string{}
	for _, p := range people {
		r = append(r, p.name)
	}
	return r
}
