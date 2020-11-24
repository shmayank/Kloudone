package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}
func main() {
	go display("Welcome")
	display("World")

}
