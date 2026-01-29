package main

import (
	"strings"
)

func cleanInput(text string) []string {
	splitText := strings.Fields(strings.ToLower(text))
	return splitText
}
