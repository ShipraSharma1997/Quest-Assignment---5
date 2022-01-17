package main

import "fmt"

var store map[string]Customer

type Customer struct {
	ID, Name, Email string
}

type CustomerStore interface {
	Create(Customer)
	Update(string, Customer) error
	Delete(string)
	GetById() (Customer, error)
	GetAll() ([]Customer, error)
}

func Create(c Customer) {

	store = map[string]Customer{
		c.ID: c,
	}
}

func Update(str string, c Customer) {
	store[str] = c

}

func main() {
	c1 := Customer{"customer_1", "name_1", "Email_1"}
	c2 := Customer{"customer_2", "name_2", "Email_2"}
	c3 := Customer{"customer_3", "name_3", "Email_3"}
	Create(c1)
	Create(c2)
	Create(c3)
	fmt.Println(store)
	Update("customer_2", c3)
	fmt.Println(store)
	delete(store, "customer_1")
	fmt.Println(store)

}
