package model

import(
	
)


type Product struct{
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	BarCode string `json:"bar"`
	Id string `json:"id"`
}

var Products  = []Product{
	{Name: "Ahomka Ginger",Description: "Ginger Flavored Candy",Price:53.80,BarCode: "6793546849",Id: "1"},
	{Name: "Coca Cola",Description: "Soda Drink, 500ml Pack of 8",Price:80.0,BarCode: "6797543459",Id: "2"},
	{Name: "Ginger Liquor",Description: "Ginger Flavored liquor",Price:33.00,BarCode: "679352424429",Id: "3"},
	{Name: "Strawberry Biscuits",Description: "Strawberry flavoured biscuits with chocolate filling",Price:85.80,BarCode: "679223149",Id: "4"},
}