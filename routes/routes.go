package routes

import (
	"github.com/gorilla/mux"
	"github.com/restapi/shop/funchandlers"
)


var AllRoutes = func (router *mux.Router)  {
	router.HandleFunc("/all/",handler.AllProducts).Methods("GET")
	router.HandleFunc("/all/{product_id}",handler.SingleProduct).Methods("GET")
	router.HandleFunc("/all/{product_id}",handler.DeleteProduct).Methods("DELETE")	
	router.HandleFunc("/all/{product_id}",handler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/all/create/",handler.CreateItem).Methods("POST")
}