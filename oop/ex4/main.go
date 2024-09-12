package main

import (
	"encoding/json"
	"fmt"
)

type Product struct{
	Name string 	`json:"name"`
	Price float64 	`json:"price"`
	Quantity int 	`json:"quantity"`
}

func Encode(p Product) string {
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	return string(data)
}

func Decode(jsonString string) Product {
	var p Product
	err := json.Unmarshal([]byte(jsonString), &p)
	if err != nil {
		fmt.Println(err)
	}
	return p
}

func main(){
	p := Product{
		Name: "Phone",
		Price: 87000.90,
		Quantity: 1,
	}

	jsonResult := Encode(p)
	fmt.Println(jsonResult)
	decodedJsonString := Decode(jsonResult)
	fmt.Println(decodedJsonString)
}