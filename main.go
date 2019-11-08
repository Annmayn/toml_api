package main

import (
	"fmt"

	"./initializer"
)

func main() {
	resourceLocation := "resource.toml"

	//returns config, map of endpoint to resource and error
	config, apiEndPoint, _ := initializer.InitializeRoutes(resourceLocation)
	fmt.Println(config)
	fmt.Println("........")
	fmt.Println(apiEndPoint)
}
