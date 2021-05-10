package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dan-dragan/go-final-web-service/controllers"
	"github.com/dan-dragan/go-final-web-service/models"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var cfg models.Config
	err := cfg.LoadConfiguration("go-final-web-service.json")
	check(err)
	controllers.RegisterControllers()
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("%s.%d", cfg.Host, cfg.Port), nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Println(fmt.Sprintf("Http server succesfully started  on %s:%d", cfg.Host, cfg.Port))
	fmt.Println("Press the Enter Key to stop anytime")
	fmt.Scanln()
}
