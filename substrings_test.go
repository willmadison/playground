package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMostRecurringSubstring(t *testing.T) {
	cases := []struct {
		given struct {
			value  string
			length int
		}
		expected string
	}{
		{
			struct {
				value  string
				length int
			}{value: "inengineering", length: 2},
			"in",
		},
		{
			struct {
				value  string
				length int
			}{value: "bananas", length: 2},
			"an",
		},
		{
			struct {
				value  string
				length int
			}{value: "bbbaaabbbaaa", length: 3},
			"bbb",
		},
		{
			struct {
				value  string
				length int
			}{value: "William Madison", length: 1},
			"i",
		},
		{
			struct {
				value  string
				length int
			}{value: "", length: 1},
			"",
		},
		{
			struct {
				value  string
				length int
			}{value: "aaabbbcccbbbaaa", length: 3},
			"bbb",
		},
		{
			struct {
				value  string
				length int
			}{value: "aabbb", length: 2},
			"bb",
		},
		{
			struct {
				value  string
				length int
			}{value: "aabbb", length: 10},
			"",
		},
		{
			struct {
				value  string
				length int
			}{value: "aabbb", length: -1},
			"",
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("FindMostRecurringSubstring(%v, %v)", tc.given.value, tc.given.length), func(t *testing.T) {
			actual := FindMostRecurringSubstring(tc.given.value, tc.given.length)
			assert.Equal(t, actual, tc.expected)
		})
	}

}
