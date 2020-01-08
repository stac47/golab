package main

import (
	"visibility/stac"
)

func main() {
	p := stac.NewPrinter("HELLO")
	p.Run()
}
