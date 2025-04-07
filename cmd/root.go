package cmd

import "fmt"

func prompt() {
	fmt.Print("> ")
}

func Execute() {
	loop()
}

func loop() {
	var name string
	for {
		prompt()
		fmt.Scan(&name)
		fmt.Printf("You entered %s\n", name)
	}
}
