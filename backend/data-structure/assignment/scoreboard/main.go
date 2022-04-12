package main

import (
	"fmt"
	"sort"
)

type Score struct {
	Name    string
	Correct int
	Wrong   int
	Empty   int
}
type Scores []Score

func (s Scores) Len() int {
	return len(s)
}

func (s Scores) Less(i, j int) bool {
	if s[i].Correct*4-s[i].Wrong > s[j].Correct*4-s[j].Wrong {
		return true
	} else if s[i].Correct*4-s[i].Wrong == s[j].Correct*4-s[j].Wrong {
		if s[i].Correct > s[j].Correct {
			return true
		} else if s[i].Correct == s[j].Correct {
			if s[i].Name < s[j].Name {
				return true
			}
		}
	}
	fmt.Println(s[i].Correct, s[j].Correct)
	return false // TODO: replace this
}

func (s Scores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Scores) TopStudents() []string {
	sort.Sort(s)
	names := []string{}
	for _, score := range s {
		names = append(names, score.Name)
	}
	return names
}

func main() {
	scores := Scores([]Score{
		{"Levi", 3, 2, 2},
		{"Agus", 5, 4, 0},
		{"Doni", 4, 0, 7},
		{"Ega", 3, 0, 7},
		{"Anton", 2, 0, 5},
	})
	sort.Sort(scores)
	for i := 0; i < scores.Len()-1; i++ {
		if scores.Less(i, i+1) {
			scores.Swap(i, i+1)
			//fmt.Println()
		}
	}
	fmt.Println(scores.TopStudents())
}
