package main

import (
	"fmt"

	"github.com/satyamkodale/call-greetings/config"
	"github.com/satyamkodale/go-greetings/v2/greetings"
)

func main() {
	fmt.Println("Version : " + config.Version)
	//calling external module
	fmt.Println(greetings.Greeting("Satyam"))
}
