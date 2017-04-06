package main

import "testing"

var (
	yesAnswer = Answer{"Yes", 1}
	noAnswer  = Answer{"No", 0}
)

func TestGenerateAnswer(t *testing.T) {
	for i := 0; i < 100000; i++ {
		a := generateAnswer()

		if *a != yesAnswer && *a != noAnswer {
			t.Errorf("returned result should be equal to a yes or no Answer, received %v", a)
		}
	}
}
