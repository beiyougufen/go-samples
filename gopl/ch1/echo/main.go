package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// 1. var + fori

	// var s, sep string
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	// fmt.Println(s)

	// 2. := + forr

	// s, sep := "", ""
	// for _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = " "
	// }
	// fmt.Println(s)

	// 3. strings.Join
	fmt.Println(strings.Join(os.Args[1:], " "))
}