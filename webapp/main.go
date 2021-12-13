package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
)

func main() {
	fmt.Println("Hello, world.")
	r := router.Generate()
	log.Fatal(http.ListenAndServe(":3000", r))
}
