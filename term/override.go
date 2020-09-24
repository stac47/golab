package main

import (
	// "fmt"
	"os"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	counter := 0
	finish := make(chan struct{})
	for {
		select {
		case t := <-tick:
			os.Stdout.Write([]byte("\r" + t.String()))
			counter++
			if counter > 10 {
				finish <- struct{}{}
			}
		case <-finish:
			return
		}
	}
}
