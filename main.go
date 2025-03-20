package main

import (
	"employee-auth/config"
	"employee-auth/routes"
	"log"
)

func main() {

	config.ConnectDatabase()

	r := routes.SetupRouter()

	if err := r.Run(":8090"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
