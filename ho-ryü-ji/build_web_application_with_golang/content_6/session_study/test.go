package main

type REST string

const (
	Get  REST = "GET"
	Post REST = "POST"
)

func RESTPatturnMatcher(expected REST, got REST) bool {
	return expected == got
}
