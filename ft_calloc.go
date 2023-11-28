package main

// ftCalloc emulates calloc behavior in C
func ftCalloc(count uintptr, size uintptr) []byte {
	sizeRes := count * size
	res := make([]byte, sizeRes)

	if len(res) != int(sizeRes) {
		return nil
	}

	ftBzero(res, int(sizeRes))

	return res
}
