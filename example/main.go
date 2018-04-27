package main

import (
	"encoding/json"
	"fmt"

	"github.com/containerum/cherry/example/exampleErrors"
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
