package main

import "valid-braces/src/utils"

type CharMap = map[int32]uint8

var BRACE_MAP = CharMap{
	'{': '}',
	'[': ']',
	'(': ')',
}

var RIGHT_BRACES = []int32{'}', ']', ')'}

func BraceYourself(braces string) bool {
	for len(braces) > 0 {
		lbs, remaining := getLeftBraces(braces)
		if len(lbs) > len(remaining) {
			return false
		}
		for i, b := range lbs {
			if remaining[len(lbs)-(i+1)] != BRACE_MAP[b] {
				return false
			}
		}
		braces = remaining[len(lbs):]
	}
	return true
}

func getLeftBraces(braces string) (lbs string, remaining string) {
	for i, b := range braces {
		if utils.Contains(RIGHT_BRACES, b) {
			lbs = braces[:i]
			remaining = braces[i:]
			return
		}
	}
	lbs = braces
	remaining = ""
	return
}
