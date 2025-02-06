package base

import "fmt"

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

type Client struct {
	Name    string
	Age     int
	Address // Embedding Composition
}

// Method
func (c *Client) ChangeCity(city string) {
	c.City = city
}

func Methods() {
	c := Client{
		Name: "John",
		Age:  25,
		Address: Address{
			Street: "123 Oak St",
			City:   "Omaha",
			State:  "NE",
			Zip:    "68106",
		},
	}
	c.ChangeCity("Chicago")
	fmt.Printf("%+v\n", c)
}
