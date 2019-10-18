package main

import (
	"log"

	loginslack "github.com/Code-Hex/login-slack"
)

func main() {
	token, err := loginslack.Login("email@email.com", "password", "workspace")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(token)
}
