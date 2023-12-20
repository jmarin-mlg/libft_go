package main

// ftStrtrim emulates strtrim behavior in C
func ftStrtrim(s1, set string) string {
	if s1 == "" || set == "" {
		return ""
	}

	start := 0
	end := ftStrlen(s1)

	for start < end && ftStrchr(set, int([]rune(s1)[start])) != -1 {
		start++
	}

	if start == end {
		trimmed, _ := ftSubstr(s1, uint(start), 0)
		return trimmed
	}

	for end > 0 && ftStrchr(set, int([]rune(s1)[end-1])) != -1 {
		end--
	}

	trimmed, _ := ftSubstr(s1, uint(start), end-start)
	return trimmed
}
