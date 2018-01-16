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

	allArray := []byte(input)

	buffer := make([]byte, 0)

	for index, ch := range allArray {
		character := string(ch)
		if strings.ToUpper(character) == character || index == len(allArray)-1 {
			if index == len(allArray)-1 {
				buffer = append(buffer, ch)
			}
			fmt.Println("A word is", string(buffer))
			buffer = nil
		}
		buffer = append(buffer, ch)
	}

}
