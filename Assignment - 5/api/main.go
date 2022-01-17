package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Customer - Our customer for all Customers
type Customer struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Customers []Customer

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllCustomers")
	json.NewEncoder(w).Encode(Customers)
}

func returnSingleCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, customer := range Customers {
		if customer.Id == key {
			json.NewEncoder(w).Encode(customer)
		}
	}
}

func createNewCustomer(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Customer struct
	// append this to our Customers array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var customer Customer
	json.Unmarshal(reqBody, &customer)
	// update our global Articles array to include
	// our new Article
	Customers = append(Customers, customer)

	json.NewEncoder(w).Encode(customer)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, customer := range Customers {
		if customer.Id == id {
			Customers = append(Customers[:index], Customers[index+1:]...)
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/customer", returnAllCustomers)
	myRouter.HandleFunc("/customer", createNewCustomer).Methods("POST")
	myRouter.HandleFunc("/customer/{id}", deleteCustomer).Methods("DELETE")
	myRouter.HandleFunc("/customer/{id}", returnSingleCustomer)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Customers = []Customer{
		Customer{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Customer{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
