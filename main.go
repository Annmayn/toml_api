package main

import (
	"fmt"
	"log"
	"net/http"

	"toml_api/handler"
	"toml_api/initializer"

	"github.com/julienschmidt/httprouter"
)

var router *httprouter.Router

func main() {
	resourceLocation := "resource.toml"

	//returns config, map of endpoint to resource and error
	config, apiEndPoint, _ := initializer.InitializeRoutes(resourceLocation)

	//do something with variables to prevent errors
	fmt.Println(config)
	fmt.Println("........")
	fmt.Println(apiEndPoint)

	router = httprouter.New()
	router.HandlerFunc("GET", "/*any", handler.CustomHandler(config, apiEndPoint))
	// router.GET("/*any", handler.CustomHandler(apiEndPoint))
	// // router.HandlerFunc("PUT", "/*any", handler.customHandler)
	// // router.HandlerFunc("POST", "/*any", handler.customHandler)
	// // router.HandlerFunc("DELETE", "/*any", handler.customHandler)
	// // router.HandlerFunc("PATCH", "/*any", handler.customHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
