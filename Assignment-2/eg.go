package main

import "fmt"

func SwapRune(r rune) rune {
	switch {
	case 97 <= r && r <= 122:
		fmt.Println("case1")
		return r - 32
	case 65 <= r && r <= 90:
		fmt.Println("case2")
		return r + 32
	default:
		fmt.Println("case3")
		return r
	}
}

func main() {
	fmt.Println(SwapRune('a'))
}
