package main



import (
	"github.com/gorilla/mux"
	"github.com/restapi/shop/routes"
	"net/http"
	"log"
)

func main()  {
	r:= mux.NewRouter()
	routes.AllRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe(":8080", r))
}