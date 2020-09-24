package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%t", os.IsExist(nil))
}
