package main

import (
	"assignment-3/config"
	"assignment-3/router"
)

func main() {
	config.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
