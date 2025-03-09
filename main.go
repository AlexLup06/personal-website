package main

import (
	"alexlupatsiy.com/personal-website/backend"
)

func main() {
	router := backend.Router()
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
