package controllers

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers(pDB *sql.DB) {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)

	pc := newProductController()

	http.Handle("/products", *pc)
	http.Handle("/products/", *pc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
