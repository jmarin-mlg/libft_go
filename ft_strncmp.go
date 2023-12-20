package main

// ftStrncmp emulates strncmp behavior in C
func ftStrncmp(s1, s2 string, n int) int {
	if n <= 0 {
		return 0
	}

	lenS1 := ftStrlen(s1)
	lenS2 := ftStrlen(s2)
	i := 0
	for i < n-1 && i < lenS1 && i < lenS2 && s1[i] == s2[i] {
		i++
	}

	if i < lenS1 && i < lenS2 {
		return int(s1[i]) - int(s2[i])
	}

	return 0
}
