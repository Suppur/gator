package main

import (
	"fmt"
	"log"

	"github.com/Suppur/gator/internal/config"
)

func main() {
	credentials, err := config.ReadConf()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Credentials: %+v\n", credentials)

	credentials.SetUser("Jamil")

	updatedCredentials, err := config.ReadConf()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated credentials: %+v\n", updatedCredentials)
}
