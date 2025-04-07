package cmd

import "fmt"

func Execute() {
	var name string

	fmt.Print("> ")
	fmt.Scan(&name)

	fmt.Printf("You entered %s", name)
}
