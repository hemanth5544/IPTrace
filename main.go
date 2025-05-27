package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var processing = false

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	mmdbConnect()

	// defer dbClose()
	// initialise()

	http.HandleFunc("GET /", getHome)
	http.HandleFunc("GET /ip/{ip}", getIp)
	http.HandleFunc("GET /random/{ipVersion}", getRandomIp)
	http.HandleFunc("GET /benchmark/{ipVersion}/{times}", getBenchmark)

	fmt.Printf("starting server on %s:%s\n", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")), nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("server closed")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
