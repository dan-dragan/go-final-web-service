package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/dan-dragan/go-final-web-service/models"
)

type productController struct {
	productIDPattern *regexp.Regexp
}

func (uc productController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/products" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.productIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		switch r.Method {
		case http.MethodGet:
			uc.get(int64(id), w)
		case http.MethodPut:
			uc.put(int64(id), w, r)
		case http.MethodDelete:
			uc.delete(int64(id), w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (uc *productController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetProducts(), w)
}

func (uc *productController) get(id int64, w http.ResponseWriter) {
	u, err := models.GetProductByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)
}

func (pc *productController) post(w http.ResponseWriter, r *http.Request) {
	p, err := pc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse product object"))
		return
	}
	p, err = models.AddProduct(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(p, w)
}

func (pc *productController) put(id int64, w http.ResponseWriter, r *http.Request) {
	p, err := pc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse product object"))
		return
	}
	if id != p.ProductId {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted product must match ID in URL"))
		return
	}
	p, err = models.UpdateProduct(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(p, w)
}

func (uc *productController) delete(id int64, w http.ResponseWriter) {
	err := models.RemoveProductById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (pc *productController) parseRequest(r *http.Request) (models.Product, error) {
	dec := json.NewDecoder(r.Body)
	var p models.Product
	err := dec.Decode(&p)
	if err != nil {
		return models.Product{}, err
	}
	return p, nil
}

func newProductController() *productController {
	return &productController{
		productIDPattern: regexp.MustCompile(`^/products/(\d+)/?`),
	}
}
