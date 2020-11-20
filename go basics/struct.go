package main

import "fmt"

func main() {
	var str string = "manish"
	for i := 0; i < len(str); i++ {
		fmt.Println(string(str[i]))
	}

}
