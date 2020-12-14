package main

import "testing"

func TestGetNumberOfQuestionsAnsweredYes(t *testing.T) {
	if GetNumberOfQuestionsAnsweredYes(`abc`) != 3 {
		t.Error("1")
	}

	if GetNumberOfQuestionsAnsweredYes(`a
a
a
a`) != 1 {
		t.Error("2")
	}
}
