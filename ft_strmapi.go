package main

// ftStrmapi emulates strmapi behavior in C
func ftStrmapi(s string, f func(uint, byte) byte) string {
	if s == "" || f == nil {
		return ""
	}

	lenS := ftStrlen(s)
	res := ftCalloc(uintptr(lenS+1), 1)
	if res == nil {
		return ""
	}

	for i := 0; i < lenS; i++ {
		res[i] = f(uint(i), s[i])
	}

	return string(res)
}
