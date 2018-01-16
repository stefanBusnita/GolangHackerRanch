package main

import (
	"fmt"
	"strings"
)

/*

 */
func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	fmt.Println("The input is: ", input)

	inputStringSlice := []byte(input)

	buffer := make([]byte, 0)

	for index, ch := range inputStringSlice {
		character := string(ch)
		if strings.ToUpper(character) == character || index == len(inputStringSlice)-1 {
			if index == len(inputStringSlice)-1 {
				buffer = append(buffer, ch)
			}
			fmt.Println(index, string(buffer))
			buffer = nil
		}
		buffer = append(buffer, ch)
	}

}
