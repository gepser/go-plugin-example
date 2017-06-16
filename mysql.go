package main

import "fmt"

type driver struct{}

//Driver database
var Driver = driver{}

func (d *driver) Init(host string) {
	fmt.Printf("[mysql]: Establishing connection to: %s\n", host)
}

//Select : simulates a Select query to the database
func (d *driver) Select(data string) {
	fmt.Printf("[mysql]: Getting some data: %s\n", data)
}

//Insert : simulates an Insertion to the database
func (d *driver) Insert(data string) {
	fmt.Printf("[mysql]: Inserting data: %s\n", data)
}
