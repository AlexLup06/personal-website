package main

import (
	"alexlupatsiy.com/personal-website/backend"
)

func main() {
	router := backend.Router()
	router.Run() // listen and serve on 0.0.0.0:8080
}
