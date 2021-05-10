package main

import (
	"log"
	"net/http"

	"github.com/dan-dragan/go-final-web-service/controllers"
)

func main() {
	controllers.RegisterControllers()
	err := http.ListenAndServe(":3000", nil)
	log.Fatal(err)
}
