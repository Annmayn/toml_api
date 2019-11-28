package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"toml_api/initializer"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/pkg/profile"
)

var router *mux.Router

//init is revoked before main()
func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env found!")
	}
}

func main() {
	defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()

	//initialize from the configuration file : resource.toml
	//returns config, map of endpoint to resource and error
	resourceLocation := "resource.toml"
	config, apiEndPoint, _ := initializer.InitializeRoutes(resourceLocation)

	//create new router
	router = mux.NewRouter()

	//todo: return error
	//initialize router to handle and validate all endpoints
	apiEndPoint = initializer.InitializeRouter(router, config, apiEndPoint)
	initializer.InitializeConfiguration(config)
	fmt.Println("Main: apiEndPoint -> ", apiEndPoint)

	// router.HandleFunc("/*any", handler.CustomHandler(config, apiEndPoint))
	// router.GET("/*any", handler.CustomHandler(apiEndPoint))
	// // router.HandlerFunc("PUT", "/*any", handler.customHandler)
	// // router.HandlerFunc("POST", "/*any", handler.customHandler)
	// // router.HandlerFunc("DELETE", "/*any", handler.customHandler)
	// // router.HandlerFunc("PATCH", "/*any", handler.customHandler)

	//TODO: Use subrouter to serve "/swaggerui/" request
	srv := &http.Server{
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Addr:         "localhost:8080",
		Handler:      initializer.RemoveTrailingSlash(router),
	}
	log.Fatal(srv.ListenAndServe())
}
