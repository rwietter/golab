package base

import "fmt"

type Client2 struct {
	Id   int
	Name string
}

// This method does not update the name of the client
// because the method receives a copy of the client
func (c Client2) UpdateName(name string) {
	c.Name = name
	fmt.Println("1. Updated name:", c.Name) // Doe
}

// This method updates the name of the client
// because the method receives a pointer to the client
// and it modifies the address of the client
func (c *Client2) UpdateNamePointer(name string) {
	c.Name = name
	fmt.Println("3. Updated name:", c.Name) // Doe
}

func NewClient(id int, name string) *Client2 {
	return &Client2{
		Id:   id,
		Name: name,
	}
}

func MutableMethod() {
	client := NewClient(1, "John")

	client.UpdateName("Doe")
	fmt.Println("2. Client name:", client.Name) // John

	client.UpdateNamePointer("Doe")
	fmt.Println("4. Client name:", client.Name) // Doe
}
