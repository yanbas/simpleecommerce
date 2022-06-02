package app

import (
	"database/sql"
	"ecommerce/model/entity"
	"ecommerce/registry"
	"ecommerce/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router     *mux.Router
	DB         *sql.DB
	ProductSvc service.ProductService
	OrderSvc   service.OrderService
}

func (a *App) Run(db *sql.DB) error {
	// register all services
	a.ProductSvc = registry.RegisterProductService(db)
	a.OrderSvc = registry.RegisterOrderService(db)

	r := mux.NewRouter()
	r.HandleFunc("/product", a.createProduct).Methods(http.MethodPost)
	r.HandleFunc("/checkout", a.checkout).Methods(http.MethodPost)
	r.HandleFunc("/product/brand", a.getProductBrand).Methods(http.MethodGet)
	r.HandleFunc("/order", a.getOrderDetail).Methods(http.MethodGet)
	r.HandleFunc("/product", a.getProducts).Methods(http.MethodGet)
	return http.ListenAndServe(":7900", r)
}

func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {
	res, err := a.ProductSvc.GetProducts()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, wrappingResponse("100", "Success", res, true))
}

func (a *App) getProductBrand(w http.ResponseWriter, r *http.Request) {
	res, err := a.ProductSvc.GetProductByBrand(r.FormValue("id"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, wrappingResponse("100", "Success", res, true))
}

func (a *App) createProduct(w http.ResponseWriter, r *http.Request) {
	var p entity.Products
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	err := a.ProductSvc.SaveProduct(p)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, wrappingResponse("000", err.Error(), nil, false))
		return
	}

	respondWithJSON(w, http.StatusCreated, wrappingResponse("100", "Success", "", true))
}

func (a *App) getOrderDetail(w http.ResponseWriter, r *http.Request) {
	res, err := a.OrderSvc.GetOrderDetail(r.FormValue("id"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, wrappingResponse("100", "Success", res, true))
}

func (a *App) checkout(w http.ResponseWriter, r *http.Request) {
	var o entity.OrderRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&o); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	err := a.OrderSvc.Checkout(o)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, wrappingResponse("000", err.Error(), nil, false))
		return
	}

	respondWithJSON(w, http.StatusCreated, wrappingResponse("100", "Success", "", true))
}

func respondWithError(w http.ResponseWriter, code int, message interface{}) {
	respondWithJSON(w, code, message)
}

func wrappingResponse(code, msg string, data interface{}, status bool) interface{} {
	result := map[string]interface{}{
		"status":  status,
		"message": msg,
		"code":    code,
		"data":    data,
	}

	return result
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
