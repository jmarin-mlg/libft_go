package st

import (
	"errors"
	"libft/src/mem"
)

// Simplified version using append()
func SimplifiedFtSplit(s string, c byte) []string {
	var split []string
	lenS := FtStrlen(s)
	iSt := 0
	iWo := -1

	for iSt <= lenS {
		if iSt == lenS || (s[iSt] == c && iWo != -1) {
			if iWo != -1 {
				split = append(split, s[iWo:iSt])
				iWo = -1
			}
		} else if s[iSt] != c && iWo == -1 {
			iWo = iSt
		}

		iSt++
	}

	return split
}

/*
** count_words: Count the number of words in a string delimited by a specified
** character.
**
** Parameters:
**   s: The input string.
**   c: Delimiting character.
**
** Returns:
**   The number of words in the string.
 */
func CountWords(s string, c byte) int {
	nWords := 0
	inWord := false

	for i := 0; i < FtStrlen(s); i++ {
		if s[i] == c {
			inWord = false
		} else if !inWord {
			inWord = true
			nWords++
		}
	}

	return nWords
}

/*
** Parameters:
**   s: The string to separate.
**   c: Delimiting character.
**
** Returns the array of new strings resulting from the split.
 */
func FtSplit(s string, c byte) []string {
	lenS := FtStrlen(s)
	nWords := CountWords(s, c)
	split := make([]string, nWords)
	iSt := 0
	iSp := 0

	for iSt < lenS {
		for iSt < lenS && s[iSt] == c {
			iSt++
		}

		if iSt < lenS {
			start := iSt

			for iSt < lenS && s[iSt] != c {
				iSt++
			}

			split[iSp] = s[start:iSt]
			iSp++
		}
	}

	return split
}

/*
** ftStrchr locates the first occurrence of c in the string s.
**
** It returns the index of the located character, or -1 if the character
** does not appear in the string.
 */
func FtStrchr(s string, c int) int {
	for i, char := range s {
		if int(char) == c {
			return i
		}
	}

	if c == 0 {
		return len(s)
	}

	return -1
}

/*
** The function computes the length of the string s. The function attempts to
** compute the length of s, but never scans beyond the first maxlen bytes of s.
**
** The function returns the number of characters that precede the terminating
** NUL character.
 */
func FtStrlen(s string) int {
	count := 0

	for range s {
		count++
	}

	return count
}

// ftStrmapi emulates strmapi behavior in C
func FtStrmapi(s string, f func(uint, byte) byte) string {
	if s == "" || f == nil {
		return ""
	}

	lenS := FtStrlen(s)
	res := mem.FtCalloc(uintptr(lenS+1), 1)
	if res == nil {
		return ""
	}

	for i := 0; i < lenS; i++ {
		res[i] = f(uint(i), s[i])
	}

	return string(res)
}

// ftStrncmp emulates strncmp behavior in C
func FtStrncmp(s1, s2 string, n int) int {
	if n <= 0 {
		return 0
	}

	lenS1 := FtStrlen(s1)
	lenS2 := FtStrlen(s2)
	i := 0
	for i < n-1 && i < lenS1 && i < lenS2 && s1[i] == s2[i] {
		i++
	}

	if i < lenS1 && i < lenS2 {
		return int(s1[i]) - int(s2[i])
	}

	return 0
}

/*
** ftStrrchr locates the last occurrence of c in the string s.
**
** It returns the index of the located character, or -1 if the character
** does not appear in the string.
 */
func FtStrrchr(s string, c int) int {
	for i := len(s) - 1; i >= 0; i-- {
		if int(s[i]) == c {
			return i
		}
	}

	if c == 0 {
		return len(s)
	}

	return -1
}

// ftStrtrim emulates strtrim behavior in C
func FtStrtrim(s1, set string) string {
	if s1 == "" || set == "" {
		return ""
	}

	start := 0
	end := FtStrlen(s1)

	for start < end && FtStrchr(set, int([]rune(s1)[start])) != -1 {
		start++
	}

	if start == end {
		trimmed, _ := FtSubstr(s1, uint(start), 0)
		return trimmed
	}

	for end > 0 && FtStrchr(set, int([]rune(s1)[end-1])) != -1 {
		end--
	}

	trimmed, _ := FtSubstr(s1, uint(start), end-start)
	return trimmed
}

/*
** ftSubstr returns a substring of s starting from index start with maximum
** length len.
**
** It returns the resulting substring or an error if memory allocation fails.
 */
func FtSubstr(s string, start uint, length int) (string, error) {
	if s == "" {
		return "", nil
	}

	size := FtStrlen(s)

	if start >= uint(size) {
		return "", nil
	}

	if length < 0 {
		length = 0
	}

	if start+uint(length) > uint(size) {
		length = size - int(start)
	}

	res := mem.FtCalloc(uintptr(length+1), 1)
	if res == nil {
		return "", errors.New("Error: memory allocation failed")
	}

	for i := 0; i < length; i++ {
		res[i] = s[start+uint(i)]
	}

	return string(res), nil
}
