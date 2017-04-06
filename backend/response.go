package main

import (
	"encoding/json"
	"math/rand"
	"time"
)

//generateAnswer is responsible for randomizing the response
func generateAnswer() *Answer {
	var stringVal = "No"
	var boolVal int

	rand.Seed(time.Now().Unix())
	num := rand.Intn(2)

	if num == 0 {
		boolVal = 1
		stringVal = "Yes"
	}

	return &Answer{stringVal, boolVal}
}

func prepareResponse() (string, error) {
	a := generateAnswer()
	str, err := a.stringInJSON()
	if err != nil {
		return "", err
	}

	return str, nil
}

func (a *Answer) stringInJSON() (string, error) {
	str, err := json.Marshal(&a)
	if err != nil {
		return "", err
	}
	return string(str), nil
}
