package main

import "valid-braces/src/utils"

type CharMap = map[uint8]uint8

var BRACE_MAP = CharMap{
	'}': '{',
	']': '[',
	')': '(',
}

var RIGHT_BRACES = []rune{'}', ']', ')'}

func BraceYourself(braces string) bool {
	from := 0
	for len(braces) > 0 {
		rb := firstRightBrace(braces, from)
		if rb <= 0 || braces[rb-1] != BRACE_MAP[braces[rb]] {
			return false
		}
		braces = braces[0:rb-1] + braces[rb+1:]
		from = rb - 1
	}
	return true
}

func firstRightBrace(braces string, start int) (rb int) {
	for i, b := range braces[start:] {
		if utils.Contains(RIGHT_BRACES, b) {
			return i + start
		}
	}
	return -1
}
