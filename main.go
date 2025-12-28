package main

import (
	"context"
	"fmt"

	"example.com/microservicechi/applications"
)

func main() {
	app := applications.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed to start server: ", err)
	}
}
