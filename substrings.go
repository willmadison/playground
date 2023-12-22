package main

import "math"

func FindMostRecurringSubstring(s string, n int) string {
	occurrencesBySubstring := map[string]int{}
	start, end := 0, 0

	highestFrequency := math.MinInt
	mostFrequent := ""

	if n > len(s) || n <= 0 {
		return ""
	}

	for end < len(s) {
		for end-start < n {
			end++
		}

		substr := s[start:end]

		occurrencesBySubstring[substr]++

		frequency := occurrencesBySubstring[substr]

		if frequency > highestFrequency {
			highestFrequency = frequency
			mostFrequent = substr
		}

		start++
	}

	return mostFrequent
}
