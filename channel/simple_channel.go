package main

import "fmt"

func main() {
	messages := make(chan string)
	messages <- "text1"
	messages <- "text2"
	fmt.Println(<-messages)
}
