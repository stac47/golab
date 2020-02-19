package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}{
		Id:   1,
		Name: "Lolo",
	}
	buf, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	fmt.Printf("%q", buf)
}
