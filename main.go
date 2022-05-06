package main

import (
	"fmt"
)

func Reverse(s string) string {
	b := []rune(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func main() {
	input := "The quick brown fox racecar jumped over the lazy dog kayak"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Printf("original:\t %q\n", input)
	fmt.Printf("reversed:\t %q\n", rev)
	fmt.Printf("reversed again:\t %q\n", doubleRev)
}
