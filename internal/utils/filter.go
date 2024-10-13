package utils

import "strings"

func processString(input string) string {
	index := strings.Index(input, "(set")
	if index == -1 {
		return input
	}

	result := input[:index]
	result = strings.TrimSpace(result)

	return result
}

func Filter(slice []string, predicate func(string) bool) []string {
	var filtered []string
	for _, str := range slice {
		if predicate(str) {
			filtered = append(filtered, processString(str))
		}
	}
	return filtered
}
