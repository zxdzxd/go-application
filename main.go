// building a simple CLI APP
package main

import (
	"fmt"
	"log"

	// "go-application/agecli"
	// wc "go-application/wordcount"
	// sf "go-application/spotifyclient"
	sc "go-application/socketcomm"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("unable to load .env file, error - %s", err)
	}
}

func main() {
	fmt.Println("Hi gopher!!")
	// agecli.GetAge()
	// wc.Wordcount()
	// sf.SfClient()
	sc.SocketClient()
	// SocketServer()
}
