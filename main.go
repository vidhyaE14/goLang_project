package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName    string  `json:"prodname"`
	ProductCategory string  `json:"prodcategory"`
	ProductPrice    float64 `json:"prodprice"`
	ProductStock    int     `json:"prodstock"`
}

var DB *gorm.DB
var err error
const dsn = "admin:Vidhya_14@tcp(sample.cbmag2acul28.us-east-1.rds.amazonaws.com:3306)/sample?charset=utf8mb4&parseTime=True&loc=Local"

func initializeRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/products",GetProducts).Methods("GET")
	router.HandleFunc("/product/{id}",GetProduct).Methods("GET")
	router.HandleFunc("/products",CreateProduct).Methods("POST")
	router.HandleFunc("/product/{id}",UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/{id}",DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func initializeMigiration(){
	DB,err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Product{})
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","applocation/json")
	var prods []Product
	DB.Find(&prods)
	json.NewEncoder(w).Encode(prods)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","applocation/json")
	params := mux.Vars(r)
	var prod Product
	DB.First(&prod, params["id"])
	json.NewEncoder(w).Encode(prod)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","applocation/json")
	var prod Product
	json.NewDecoder(r.Body).Decode(&prod)
	DB.Create(&prod)
	json.NewEncoder(w).Encode(prod)

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","applocation/json")
	params := mux.Vars(r)
	var prod Product
	DB.First(&prod, params["id"])
	json.NewDecoder(r.Body).Decode(&prod)
	DB.Save(&prod)
	json.NewEncoder(w).Encode(prod)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","applocation/json")
	params := mux.Vars(r)
	var prod Product
	DB.Delete(&prod, params["id"])
	json.NewEncoder(w).Encode("The product is deleted succesfully!!!")
}

func main() {
	initializeMigiration()
	initializeRouter()
}