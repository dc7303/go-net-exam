package main

import (
	"fmt"
	"log"

	"go-net-exam/app"
)

func main() {
	if err := app.Run(":8080"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Run server. http://localhost:8080")
	}
}
