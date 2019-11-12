package main

import (
	"fmt"
	"log"
	"net/http"
	"toml_api/initializer"
	"github.com/gorilla/mux"
)

var router *mux.Router

func main() {
	resourceLocation := "resource.toml"

	//initialize from the configuration file : resource.toml
	//returns config, map of endpoint to resource and error
	config, apiEndPoint, _ := initializer.InitializeRoutes(resourceLocation)

	//create new router
	router = mux.NewRouter()

	//todo: return error
	//initialize router to handle and validate all endpoints
	apiEndPoint = initializer.InitializeRouter(router, config, apiEndPoint)
	fmt.Println("Main: apiEndPoint -> ", apiEndPoint)

	// router.HandleFunc("/*any", handler.CustomHandler(config, apiEndPoint))
	// router.GET("/*any", handler.CustomHandler(apiEndPoint))
	// // router.HandlerFunc("PUT", "/*any", handler.customHandler)
	// // router.HandlerFunc("POST", "/*any", handler.customHandler)
	// // router.HandlerFunc("DELETE", "/*any", handler.customHandler)
	// // router.HandlerFunc("PATCH", "/*any", handler.customHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
