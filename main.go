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
