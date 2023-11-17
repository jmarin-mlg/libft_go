package main

import (
	"fmt"
)

func ftSplit(s string, c byte) []string {
	lenS := len(s)
	iSt := 0
	iWo := -42
	var split []string

	for iSt <= lenS {
		if iSt == lenS || (s[iSt] == c && iWo != -42) {
			if iWo != -42 {
				word := s[iWo:iSt]
				split = append(split, word)
				iWo = -42
			}
		} else if s[iSt] != c && iWo == -42 {
			iWo = iSt
		}

		iSt++
	}

	return split
}

func main() {
	s := "This is a test"
	c := byte(' ')

	result := ftSplit(s, c)
	fmt.Println(result)
}
