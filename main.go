package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/dan-dragan/go-final-web-service/controllers"
	"github.com/dan-dragan/go-final-web-service/models"
	_ "github.com/go-sql-driver/mysql"
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

	pDb, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/test", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port))
	check(err)
	defer pDb.Close()

	controllers.RegisterControllers(pDb)

	go func() {
		err := http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Println(fmt.Sprintf("Http server succesfully started  on %s:%d", cfg.Host, cfg.Port))
	fmt.Println("Press the Enter Key to stop anytime")
	fmt.Scanln()
}
