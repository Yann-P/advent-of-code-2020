package main

import "strings"

func GetNumberOfQuestionsAnsweredYesByEveryone(group string) int {

	persons := strings.Split(group, "\n")
	setOfQuestions := make(map[string]bool)
	questionsAnsweredYesCount := make(map[string]int)
	nPersons := 0

	for _, person := range persons {

		if person == "" {
			continue
		}

		nPersons++

		questions := strings.Split(person, "")

		uniqueQuestionsFromPerson := make(map[string]bool)

		for _, question := range questions {
			uniqueQuestionsFromPerson[question] = true
		}

		for question := range uniqueQuestionsFromPerson {
			setOfQuestions[question] = true
			questionsAnsweredYesCount[question]++
		}

	}

	questionsEveryoneAnsweredYesTo := 0

	for question := range setOfQuestions {
		if questionsAnsweredYesCount[question] == nPersons {
			questionsEveryoneAnsweredYesTo++
		}
	}

	return questionsEveryoneAnsweredYesTo
}
