package main

// ftStrncmp emulates strncmp behavior in C
func ftStrncmp(s1, s2 string, n int) int {
	if n <= 0 {
		return 0
	}

	i := 0
	for i < n-1 && i < len(s1) && i < len(s2) && s1[i] == s2[i] {
		i++
	}

	if i < len(s1) && i < len(s2) {
		return int(s1[i]) - int(s2[i])
	}

	return 0
}
