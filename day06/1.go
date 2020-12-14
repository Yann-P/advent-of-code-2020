package main

import "strings"

func GetNumberOfQuestionsAnsweredYes(group string) int {

	persons := strings.Split(group, "\n")
	questionsAnsweredYes := make(map[string]bool)

	for _, person := range persons {
		questions := strings.Split(person, "")

		for _, question := range questions {
			questionsAnsweredYes[question] = true
		}

	}

	return len(questionsAnsweredYes)
}
