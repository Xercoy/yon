package main

// Answer is used to represent a yes or no for the API. It's Json unmarshaling
// is what is going to be used as a response.
type Answer struct {
	StringVal string `json:"string,"`
	BoolVal   int    `json:"boolean,"`
}
