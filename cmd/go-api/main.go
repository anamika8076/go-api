package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anamika8076/go-api/internal/config"
	"github.com/anamika8076/go-api/internal/handlers/student"
)




func main(){
	
	// Load the configuration
	cfg := config.MustLoad()

	//database setup
	//setup router
	router :=http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New())


	//setup server
	server := http.Server{
		Addr: cfg.Addr,
		Handler: router,
	}
	fmt.Println("server started")
	go func(){
	    err := server.ListenAndServe()
	    if err != nil {
		    log.Fatal("failed to start sever")

	} 


	} ()
	




}