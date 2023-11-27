package main

// Simplified version using append()
func simplifiedFtSplit(s string, c byte) []string {
	var split []string
	lenS := ftStrlen(s)
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
