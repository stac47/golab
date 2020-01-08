package stac

import "fmt"

type printer struct {
	Message string
}

func NewPrinter(message string) printer {
	return printer{
		message,
	}
}

func (p *printer) Run() {
	fmt.Printf("Message: %s\n", p.Message)
}
