package main

import "fmt"

func main() {
	var arr [3]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 5
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
