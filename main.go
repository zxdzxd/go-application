// building a simple CLI APP
package main

import "fmt"

func main() {
	fmt.Println("Hi gopher!!")
	fmt.Print("enter your Name")
	var name string
	fmt.Scanln(&name)
}

