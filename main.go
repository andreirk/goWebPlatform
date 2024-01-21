package main

import (
	"webPlatform/placeholder"
	"webPlatform/services"
)

func main() {
	services.RegisterDefaultServices()
	placeholder.Start()
}
