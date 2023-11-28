package main

import (
	"fmt"
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
}
