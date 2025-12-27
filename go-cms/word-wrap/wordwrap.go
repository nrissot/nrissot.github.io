package wordwrap

import (
	"strings"
)

// STATEMENT  : This library aims for a good enough result, not a perfect result
// 				Text wrapping is not a matter of perfection.

func Wrap(s string, w int) string {
	var out string = ""
	var line_length = 0
	for _, words := range strings.Split(s, " ") {
		for i, word := range strings.Split(words, "\n") {
			if i != 0 {
				out += "\n"
				line_length = 0
			}
			word_length := len([]rune(word))
			if line_length == 0 {
				out += word
				line_length += word_length
			} else if line_length+word_length+1 >= w {
				out += "\n" + word
				line_length = word_length
			} else {
				out += " " + word
				line_length += word_length + 1
			}
		}
	}
	return out
}

func WrapBreak(s string, w int) string {
	var out string = ""
	var line_length = 0
	for _, word := range strings.Split(s, " ") {
		word_length := len([]rune(word))
		if line_length+word_length+1 <= w {
			if line_length != 0 {
				out += " "
				line_length += 1
			}
			out += word
			line_length += word_length
		} else {
			after := (line_length + word_length) - w
			before := word_length - after - 1
			if before >= word_length/2 {
				out += " " + word[:before-1] + "-\n-" + word[before-1:]
				line_length = after + 3
			} else {
				out += "\n" + word
				line_length = word_length
			}

		}
	}
	return out
}
