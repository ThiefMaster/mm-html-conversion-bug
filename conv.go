package main

import (
	"fmt"
	"os"
	"strings"

	"code.sajari.com/docconv/v2"
)

func main() {
	data, err := os.ReadFile("message.html")
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	htmlMessage := string(data)
	text, _, err := docconv.ConvertHTML(strings.NewReader(htmlMessage), true)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Printf("result: %s\n", text)
	}
}
