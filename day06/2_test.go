package main

import "testing"

func TestGetNumberOfQuestionsAnsweredYesByEveryone(t *testing.T) {
	if GetNumberOfQuestionsAnsweredYesByEveryone(`abc`) != 3 {
		t.Error("1")
	}

	if GetNumberOfQuestionsAnsweredYesByEveryone(`a
b
c`) != 0 {
		t.Error("2")
	}

	if GetNumberOfQuestionsAnsweredYesByEveryone(`a
a
a
a`) != 1 {
		t.Error("3")
	}

	if GetNumberOfQuestionsAnsweredYesByEveryone(`ab
ac`) != 1 {
		t.Error("4")
	}

	if GetNumberOfQuestionsAnsweredYesByEveryone(`b`) != 1 {
		t.Error("5")
	}

	if GetNumberOfQuestionsAnsweredYesByEveryone(`adc
dbznn
mdn`) != 1 {
		t.Error("6")
	}

	if GetNumberOfQuestionsAnsweredYesByEveryone(`ab
ba
ab
ba`) != 2 {
		t.Error("7")
	}

	if GetNumberOfQuestionsAnsweredYesByEveryone(`abcdef
ecba
`) != 4 {
		t.Error(GetNumberOfQuestionsAnsweredYesByEveryone(`abcdef
ecba
`))
	}
}
