package main

import (
	"encoding/json"
	"fmt"

	"git.containerum.net/ch/cherry/example/exampleErrors"
)

//go:generate noice -t ./errors.toml

func main() {
	data, err := json.MarshalIndent(exampleErrors.ErrInvalidCheese(), "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
