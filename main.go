package main

import (
	"log"

	"alexlupatsiy.com/personal-website/backend"
)

func main() {
	if err := backend.Router(); err != nil {
		log.Fatal(err)
	}
}
