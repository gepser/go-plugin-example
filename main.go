package main

import (
	"fmt"
	"os"
	"plugin"
)

func main() {

	p, err := plugin.Open("plugins/mysql.so")

	if err != nil {
		fmt.Print(fmt.Errorf("Trying to Open a plugin: %v", err))
		os.Exit(1)
	}

	dbDriver, err := p.Lookup("Driver")

	if err != nil {
		fmt.Print(fmt.Errorf("Lookup the Insert symbol: %v", err))
		os.Exit(1)
	}

	// Casting here so I don't have to do it everywhere
	// And yes, we need to validate this cast

	d := dbDriver.(*driver)

	// Initializing the driver (yeah, we need a constructor instead)
	// And just in case you are wondering why I didn't use the 'init' function
	// it's because it gets executed twice while Opening, for reasons

	d.Init("http://localhost:3306/")

	// Here we are going to simulate an insertion to the database

	d.Insert("Golang")

	// Now we are going to simulate a select query to the database

	d.Select("Gepser")
}
