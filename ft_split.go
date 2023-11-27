package main

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
func countWords(s string, c byte) int {
	nWords := 0
	inWord := false

	for i := 0; i < ftStrlen(s); i++ {
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
func ftSplit(s string, c byte) []string {
	lenS := ftStrlen(s)
	nWords := countWords(s, c)
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
