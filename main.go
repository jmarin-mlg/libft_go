package main

import (
	"fmt"
	"libft/src/is"
	"libft/src/mem"
	"libft/src/put"
	"libft/src/st"
	"libft/src/to"
	"os"
	"strconv"
	"strings"
)

func main() {
	// ft_atoi
	str := "-123"
	resultOriginalAtoi, err := strconv.Atoi(str)
	resultFtAtoi := to.FtAtoi(str)

	fmt.Printf("\n### ATOI ###\n\n")
	fmt.Printf("        Original String: \"%s\"\n", str)

	if err != nil {
		fmt.Printf("            Atoi result: %d - Error: %s\n", resultOriginalAtoi, err)
	} else {
		fmt.Printf("            Atoi result: %d\n", resultOriginalAtoi)
	}

	fmt.Printf("          FtAtoi result: %d\n", resultFtAtoi)

	// ft_memset
	buffer := make([]byte, 10)

	fmt.Printf("\n### MEMSET ###\n\n")
	mem.FtMemset(buffer, 'A', len(buffer))
	fmt.Println(string(buffer))

	// ft_bzero
	fmt.Printf("\n### BZERO ###\n\n")
	mem.FtBzero(buffer, 5)
	fmt.Println(string(buffer))

	// ft_calloc
	count, size := uintptr(5), uintptr(4)
	memory := mem.FtCalloc(count, size)

	fmt.Printf("\n### CALLOC ###\n\n")

	if memory == nil {
		fmt.Println("Error: memory allocation failed")
	} else {
		for i := uintptr(0); i < count*size; i++ {
			fmt.Printf("%d ", memory[i])
		}

		fmt.Println()
	}

	// ft_memchr, ft_memcmp, ft_memcpy, ft_memmove
	fmt.Printf("\n### MEMCHR MEMCMP MEMCPY MEMMOVE ###\n\n")

	// ftMemchr
	str1 := "Hello, World!"
	charToFind := byte('o')
	n1 := 5
	resultMemchr := mem.FtMemchr([]byte(str1), int(charToFind), n1)
	fmt.Printf(" ftMemchr Result: %v\n", resultMemchr)

	// ftMemcmp
	str2a := "Hello, Go!"
	str2b := "Hello, World!"
	n2 := 6
	resultMemcmp := mem.FtMemcmp([]byte(str2a), []byte(str2b), n2)
	fmt.Printf(" ftMemcmp Result: %v\n", resultMemcmp)

	// ftMemcpy
	src := "Copy this!"
	dst := make([]byte, len(src))
	mem.FtMemcpy(dst, []byte(src), len(src))
	fmt.Printf(" ftMemcpy Result: %s\n", dst)

	// ftMemmove
	str3 := "Move me!"
	dst3 := make([]byte, len(str3))
	src3 := []byte(str3)
	len3 := len(str3)
	mem.FtMemmove(dst3, src3, len3)
	fmt.Printf("ftMemmove Result: %s\n", dst3)

	// ft_isalnum, ft_isalpha, ft_isascii, ft_isdigit, ft_isprint
	char1 := 'A'
	char2 := '5'
	char3 := 'ç'
	char4 := 'X'
	char5 := '7'

	fmt.Printf("\n### ISPRINT ISDIGIT ISASCII ISALPHA ISALNUM ###\n\n")
	fmt.Printf("Is '%c' printable? %t\n", char1, is.FtIsPrint(int(char1)))
	fmt.Printf("Is '%c' a digit? %t\n", char2, is.FtIsDigit(int(char2)))
	fmt.Printf("Is '%c' an ASCII character? %t\n", char3, is.FtIsASCII(int(char3)))
	fmt.Printf("Is '%c' an alphabetic character? %t\n", char4, is.FtIsAlpha(int(char4)))
	fmt.Printf("Is '%c' alphanumeric? %t\n", char5, is.FtIsAlnum(int(char5)))

	// ft_itoa
	num := -123
	resultOriginalItoa := strconv.Itoa(num)
	resultFtItoa, err := to.FtItoa(num)

	fmt.Printf("\n### ITOA ###\n\n")
	fmt.Printf("Original number: %d\n", num)
	fmt.Printf("    Itoa result: \"%s\"\n", resultOriginalItoa)

	if err != nil {
		fmt.Printf("  ftItoa result: \"%s\" - Error: %s\n", resultFtItoa, err)
	} else {
		fmt.Printf("  ftItoa result: \"%s\"\n", resultFtItoa)
	}

	// ft_putchar_fd, ft_putendl_fd, ft_putnbr_fd, ft_putstr_fd
	fileCreated, errFileCreated := os.Create("output.txt")

	if errFileCreated != nil {
		fmt.Println("Error creating file:", errFileCreated)
		return
	}

	defer fileCreated.Close()

	fmt.Printf("\n### PUTCHAR PUTSTR PUTNBR PUTENDL ###\n\n")
	put.FtPutcharFd('A', fileCreated)
	put.FtPutendlFd("Hello, World!", fileCreated)
	put.FtPutnbrFd(12345, fileCreated)
	put.FtPutstrFd("This is a string.", fileCreated)

	fileRead, errFileRead := os.ReadFile("output.txt")

	if errFileRead != nil {
		fmt.Println("Error al leer el archivo:", errFileRead)
		return
	}

	fmt.Println("Content file:")
	fmt.Println(string(fileRead))

	err = os.Remove("output.txt")

	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}

	// ft_split
	str = "This,is,a,test"
	delimiter := ","

	resultOriginalSplit := strings.Split(str, delimiter)
	resultFtSplit := st.FtSplit(str, []byte(delimiter)[0])
	resultSimplifiedFtSplit := st.SimplifiedFtSplit(str, []byte(delimiter)[0])

	fmt.Printf("\n### SPLIT ###\n\n")
	fmt.Printf("          Original String: %s\n", str)
	fmt.Printf("             Split result: %v\n", resultOriginalSplit)
	fmt.Printf("           ftSplit result: %v\n", resultFtSplit)
	fmt.Printf("Simplified ftSplit result: %v\n", resultSimplifiedFtSplit)

	// ft_strchr, ft_strrchr
	fmt.Printf("\n### STRCHR STRRCHR ###\n\n")

	// Prueba con ftStrchr
	s1 := "Hello, World!"
	charTolower := byte('o')
	index1 := st.FtStrchr(s1, int(charTolower))

	if index1 != -1 {
		fmt.Printf("Encontrado '%c' en la posición %d de \"%s\"\n", charTolower, index1, s1)
	} else {
		fmt.Printf("'%c' no encontrado en \"%s\"\n", charTolower, s1)
	}

	// Prueba con ftStrrchr
	s2 := "Go Programming"
	charToupper := byte('o')
	index2 := st.FtStrrchr(s2, int(charToupper))

	if index2 != -1 {
		fmt.Printf("Encontrado '%c' en la posición %d de \"%s\"\n", charToupper, index2, s2)
	} else {
		fmt.Printf("'%c' no encontrado en \"%s\"\n", charToupper, s2)
	}

	// ft_strmapi
	fmt.Printf("\n### STRMAPI ###\n\n")

	// Prueba con ftStrmapi
	s5 := "Hello, World!"

	// Una función de ejemplo que convierte letras minúsculas en mayúsculas y viceversa
	convertir := func(index uint, char byte) byte {
		if char >= 'a' && char <= 'z' {
			return char - ('a' - 'A')
		} else if char >= 'A' && char <= 'Z' {
			return char + ('a' - 'A')
		}
		return char
	}

	resultFtStrmapi := st.FtStrmapi(s5, convertir)

	fmt.Printf("Original string: \"%s\"\n", s5)
	fmt.Printf("Result after conversion: \"%s\"\n", resultFtStrmapi)

	// ft_strncmp
	fmt.Printf("\n### STRNCMP ###\n\n")

	// Prueba con ftStrncmp
	s3 := "Hello, World!"
	s4 := "Hello, Go!"
	n := 7

	resultFtStrcmp := st.FtStrncmp(s3, s4, n)

	fmt.Printf("s1: \"%s\"\n", s3)
	fmt.Printf("s2: \"%s\"\n", s4)
	fmt.Printf("Comparison result (first %d characters): %d\n", n, resultFtStrcmp)

	// ft_strtrim
	fmt.Printf("\n### STRTRIM ###\n\n")

	// Prueba con ftStrtrim
	s := "   Hello, World!   "
	set := " "
	trimmed := st.FtStrtrim(s, set)

	fmt.Printf("Original: \"%s\"\n", s)
	fmt.Printf("Set: \"%s\"\n", set)
	fmt.Printf("Trimmed: \"%s\"\n", trimmed)

	// ft_substr
	inputString := "Hello, World!"
	startIndex := uint(7)
	substrLength := 5

	fmt.Printf("\n### SUBSTR ###\n\n")

	result, err := st.FtSubstr(inputString, startIndex, substrLength)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf(" Original string: %s\n", inputString)
	fmt.Printf("Substring result: %s\n", result)

	// ft_tolower, ft_toupper
	characters := []int{'A', 'b', 'C', 'd', 'E', '!', '1'}

	fmt.Printf("\n### TOLOWER TOUPPER ###\n\n")
	fmt.Println("Original characters:", StringSlice(characters))

	// Test ftTolower
	for i, c := range characters {
		characters[i] = to.FtTolower(c)
	}
	fmt.Println("After ftTolower:", StringSlice(characters))

	// Test ftToupper
	for i, c := range characters {
		characters[i] = to.FtToupper(c)
	}
	fmt.Println("After ftToupper:", StringSlice(characters))
}

func StringSlice(intSlice []int) string {
	var result string

	for _, val := range intSlice {
		result += string(val)
	}

	return result
}
