package main

func ftStrlen(s string) int {
	count := 0

	for range s {
		count++
	}

	return count
}