package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dan-dragan/go-final-web-service/controllers"
)

func main() {
	controllers.RegisterControllers()
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Http server succesfully started  on port 3000")
	}

}
