package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	// ft_atoi
	str := "-123"
	resultOriginalAtoi, err := strconv.Atoi(str)
	resultFtAtoi := ftAtoi(str)

	fmt.Printf("\n### ATOI ###\n\n")
	fmt.Printf("        Original String: \"%s\"\n", str)

	if err != nil {
		fmt.Printf("            Atoi result: %d - Error: %s\n", resultOriginalAtoi, err)
	} else {
		fmt.Printf("            Atoi result: %d\n", resultOriginalAtoi)
	}

	fmt.Printf("          ftAtoi result: %d\n", resultFtAtoi)

	// ft_memset
	buffer := make([]byte, 10)

	fmt.Printf("\n### MEMSET ###\n\n")
	ftMemset(buffer, 'A', len(buffer))
	fmt.Println(string(buffer))

	// ft_bzero
	fmt.Printf("\n### BZERO ###\n\n")
	ftBzero(buffer, 5)
	fmt.Println(string(buffer))

	// ft_calloc
	count, size := uintptr(5), uintptr(4)
	mem := ftCalloc(count, size)

	fmt.Printf("\n### CALLOC ###\n\n")

	if mem == nil {
		fmt.Println("Error: memory allocation failed")
	} else {
		for i := uintptr(0); i < count*size; i++ {
			fmt.Printf("%d ", mem[i])
		}

		fmt.Println()
	}

	// ft_isalnum, ft_isalpha, ft_isascii, ft_isdigit, ft_isprint
	char1 := 'A'
	char2 := '5'
	char3 := 'ç'
	char4 := 'X'
	char5 := '7'

	fmt.Printf("\n### ISPRINT ISDIGIT ISASCII ISALPHA ISALNUM ###\n\n")
	fmt.Printf("Is '%c' printable? %t\n", char1, ftIsPrint(int(char1)))
	fmt.Printf("Is '%c' a digit? %t\n", char2, ftIsDigit(int(char2)))
	fmt.Printf("Is '%c' an ASCII character? %t\n", char3, ftIsASCII(int(char3)))
	fmt.Printf("Is '%c' an alphabetic character? %t\n", char4, ftIsAlpha(int(char4)))
	fmt.Printf("Is '%c' alphanumeric? %t\n", char5, ftIsAlnum(int(char5)))

	// ft_itoa
	num := -123
	resultOriginalItoa := strconv.Itoa(num)
	resultFtItoa, err := ftItoa(num)

	fmt.Printf("\n### ITOA ###\n\n")
	fmt.Printf("Original number: %d\n", num)
	fmt.Printf("    Itoa result: \"%s\"\n", resultOriginalItoa)

	if err != nil {
		fmt.Printf("  ftItoa result: \"%s\" - Error: %s\n", resultFtItoa, err)
	} else {
		fmt.Printf("  ftItoa result: \"%s\"\n", resultFtItoa)
	}

	// ft_memchr, ft_memcmp, ft_memcpy, ft_memmove
	fmt.Printf("\n### MEMCHR MEMCMP MEMCPY MEMMOVE ###\n\n")

	// ftMemchr
	str1 := "Hello, World!"
	charToFind := byte('o')
	n1 := 5
	resultMemchr := ftMemchr([]byte(str1), int(charToFind), n1)
	fmt.Printf(" ftMemchr Result: %v\n", resultMemchr)

	// ftMemcmp
	str2a := "Hello, Go!"
	str2b := "Hello, World!"
	n2 := 6
	resultMemcmp := ftMemcmp([]byte(str2a), []byte(str2b), n2)
	fmt.Printf(" ftMemcmp Result: %v\n", resultMemcmp)

	// ftMemcpy
	src := "Copy this!"
	dst := make([]byte, len(src))
	ftMemcpy(dst, []byte(src), len(src))
	fmt.Printf(" ftMemcpy Result: %s\n", dst)

	// ftMemmove
	str3 := "Move me!"
	dst3 := make([]byte, len(str3))
	src3 := []byte(str3)
	len3 := len(str3)
	ftMemmove(dst3, src3, len3)
	fmt.Printf("ftMemmove Result: %s\n", dst3)

	// ft_putchar_fd, ft_putendl_fd, ft_putnbr_fd, ft_putstr_fd
	fileCreated, errFileCreated := os.Create("output.txt")

	if errFileCreated != nil {
		fmt.Println("Error creating file:", errFileCreated)
		return
	}

	defer fileCreated.Close()

	fmt.Printf("\n### PUTCHAR PUTSTR PUTNBR PUTENDL ###\n\n")
	ftPutcharFd('A', fileCreated)
	ftPutendlFd("Hello, World!", fileCreated)
	ftPutnbrFd(12345, fileCreated)
	ftPutstrFd("This is a string.", fileCreated)

	fileRead, errFileRead := ioutil.ReadFile("output.txt")

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
	resultFtSplit := ftSplit(str, []byte(delimiter)[0])
	resultSimplifiedFtSplit := simplifiedFtSplit(str, []byte(delimiter)[0])

	fmt.Printf("\n### SPLIT ###\n\n")
	fmt.Printf("          Original String: %s\n", str)
	fmt.Printf("             Split result: %v\n", resultOriginalSplit)
	fmt.Printf("           ftSplit result: %v\n", resultFtSplit)
	fmt.Printf("Simplified ftSplit result: %v\n", resultSimplifiedFtSplit)

	// ft_substr
	inputString := "Hello, World!"
	startIndex := uint(7)
	substrLength := 5

	fmt.Printf("\n### SUBSTR ###\n\n")

	result, err := ftSubstr(inputString, startIndex, substrLength)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf(" Original string: %s\n", inputString)
	fmt.Printf("Substring result: %s\n", result)

	// ft_tolower, ft_toupper
	characters := []int{'A', 'b', 'C', 'd', 'E', '!', '1'}

	fmt.Printf("\n### TOLOWER TOUPPER ###\n\n")
	fmt.Println("Original characters:", stringSlice(characters))

	// Test ftTolower
	for i, c := range characters {
		characters[i] = ftTolower(c)
	}
	fmt.Println("After ftTolower:", stringSlice(characters))

	// Test ftToupper
	for i, c := range characters {
		characters[i] = ftToupper(c)
	}
	fmt.Println("After ftToupper:", stringSlice(characters))
}

func stringSlice(intSlice []int) string {
	var result string

	for _, val := range intSlice {
		result += string(val)
	}

	return result
}
