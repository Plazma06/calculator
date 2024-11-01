package main

import (
	"fmt"
	"strings"
)

func main() {
	res := strings.Join(strings.Fields("a  b  c"), "")
	fmt.Println(len(res), res)

}

func deleteDuplicates(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] != s[i-1] {
			res += string(s[i])
		}
	}
	return res
}
