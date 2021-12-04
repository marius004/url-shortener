package main

import (
	"fmt"
	"log"
)

func main() {
	config := NewConfig()
	_, err := Connect(CreateDatabaseDSN(config))

	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("It Worked")
	}
}
