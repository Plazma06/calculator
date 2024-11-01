package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"

	// "strconv"
	"strings"
)

func main() {
	fmt.Print("Input: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	input = strings.Join(strings.Fields(input), "")

	var num string
	var oper rune = ' '
	var res int
	//var operations = make(map[rune]int)

	for i, c := range input {
		if unicode.IsDigit(c) {
			if i == len(input) - 1 {
				num += string(c)
			} else {
				num += string(c)
				continue
			}
		} else if c == rune(input[i+1]) {
			panic("Error Pon")
		}
		
		if oper == ' ' {
			res, _ = strconv.Atoi(num)
		} 

		n, _ := strconv.Atoi(num)
		switch oper {
		case '+':
			res += n
		case '-':
			res -= n
		}

		oper = c
		num = ""
	}

	fmt.Println("Output:", input, "=", res)
}

