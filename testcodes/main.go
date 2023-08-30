package main

import (
	"fmt"
	"strings"
)



func StringToLLVM(input string) (string, int) {
	// Replace %d with %ld
	transformed := strings.Replace(input, "%d", "%ld", -1)

	// Replace \n with \0A
	transformed = strings.Replace(transformed, "\n", "\\0A", -1)

	// Add \00 at the end
	transformed += "\\00"

	// Compute character count
	charCount := 0
	
	newLineCount := strings.Count(transformed, "\\0A") 

	// charcount = (len(string) - 3 * newLineCount - 3) + newLineCount + 1
	charCount = len(transformed) - 2 * (newLineCount + 1)

	return transformed, charCount
}


func main() {
	input := "offset=%d\npt2.x=%d\npt2.y=%d\n"

	transformed, charCount := StringToLLVM(input)
	fmt.Println(transformed)
	fmt.Println(charCount)

	// offset=%ld\0Apt2.x=%ld\0Apt2.y=%ld\0A\00
	// offset=%ld\0Apt2.x=%ld\0Apt2.y=%ld\0A\00
}