package main

import (
	"fmt"
	"os"
	"plugin"
)

func main() {

	p, err := plugin.Open("./add.so")

	if err != nil {
		fmt.Print(fmt.Errorf("Trying to Open a plugin: %v", err))
		os.Exit(1)
	}

	operation, err := p.Lookup("Add")

	if err != nil {
		fmt.Print(fmt.Errorf("Lookup the Add symbol: %v", err))
		os.Exit(1)
	}

	result := operation.(func(int, int) int)(1, 2)
	fmt.Println(result)

}
