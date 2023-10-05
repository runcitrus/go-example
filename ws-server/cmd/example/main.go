package main

import (
	"fmt"
	"log"
	"os"

	"example/internal/app"
)

func main() {
	arg := ""
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}

	switch arg {
	case "-v", "--version":
		fmt.Println("Example App v0.0.1")
	default:
		log.Println("Starting...")
		if err := app.Start(); err != nil {
			log.Fatal(err)
		}
	}
}
