package main

import (
	"net/http"

	"github.com/dan-dragan/go-learn-web-service/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
