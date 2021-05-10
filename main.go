package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dan-dragan/go-final-web-service/controllers"
)

func main() {
	controllers.RegisterControllers()
	go func() {
		err := http.ListenAndServe(":3000", nil)
		if err != nil {
			log.Fatal(err)
		} 
	}
	fmt.Println("Http server succesfully started  on port 3000")
	fmt.Println("Press the Enter Key to stop anytime")
	fmt.Scanln()
}
