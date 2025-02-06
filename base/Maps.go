package base

import "fmt"

func Maps() {
	salary := map[string]float64{
		"steve": 12000.50,
		"jamie": 15000.0,
		"mike":  9000.0,
	}

	// Get the value of a specified key
	fmt.Println("Salary of steve is", salary["steve"])

	// Set the value of a specified key
	salary["mike"] = 18000.0

	// Check if a specified key is present in the map
	value, ok := salary["jamie"]
	if ok {
		fmt.Println("Salary of jamie is", value)
	} else {
		fmt.Println("Salary of jamie is not present")
	}

	// Delete the key/value pair for a specified key
	delete(salary, "steve")

	// Loop through all key/value pairs in a map
	for key, value := range salary {
		fmt.Printf("Salary of %s is %f\n", key, value)
	}

	// Create a new map
	employee := make(map[string]float64)
	employee["steve"] = 12000.0
	fmt.Printf("Employee map: %v\n", employee)

	// Maps of maps
	user4 := map[string]map[string]string{
		"user1": {
			"name":     "John",
			"lastname": "Doe",
		},
		"user2": {
			"name":     "Jane",
			"lastname": "Doe",
		},
	}
	fmt.Println(user4)
}
