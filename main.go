package main

import (
	"log"

	backend "alexlupatsiy.com/personal-website/internal"
)

func main() {
	if err := backend.Router(); err != nil {
		log.Fatal(err)
	}
}
