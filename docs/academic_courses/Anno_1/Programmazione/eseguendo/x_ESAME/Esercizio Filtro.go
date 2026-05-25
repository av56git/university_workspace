package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := []rune(os.Args[1])
	k, _ := strconv.Atoi(os.Args[2])
	for i := 0; i < len(s); i++ {
		s[i] = s[i] + rune(k)
	}
	fmt.Println("stringa cifrata: ", string(s))
}
