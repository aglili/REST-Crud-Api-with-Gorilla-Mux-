package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/restapi/shop/models"
)



func AllProducts(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.Products)
}

func SingleProduct(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r)
	for _,item := range model.Products{
		if item.Id == parameters["product_id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}


func CreateItem(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var item model.Product
	_ = json.NewDecoder(r.Body).Decode(&item)
	nextID := len(model.Products) +1
	item.Id = strconv.Itoa(nextID)
	model.Products = append(model.Products, item)
	json.NewEncoder(w).Encode(model.Products)
}


func DeleteProduct(w http.ResponseWriter, r *http.Request){
	parameters := mux.Vars(r)
	id := parameters["product_id"]
	for index,item := range model.Products{
		if item.Id == id{
			model.Products = append(model.Products[:index], model.Products[index+1:]...)
			break
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Products)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request){
	parameters := mux.Vars(r)
	for index,item := range model.Products{
		if item.Id == parameters["product_id"]{
			model.Products = append(model.Products[:index], model.Products[index+1:]...)
			var item model.Product
			_ = json.NewDecoder(r.Body).Decode(&item)
			item.Id = parameters["product_id"]
			model.Products = append(model.Products, item)
			json.NewEncoder(w).Encode(model.Products)
			return
		}
	}
	json.NewEncoder(w).Encode(model.Products)
}