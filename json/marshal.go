package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := Person{
		Name: "lolo",
		Age:  12,
	}
	buf, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("%s\n", buf)
	bb := bytes.NewBuffer(buf)
	io.Copy(os.Stdout, bb)
}
