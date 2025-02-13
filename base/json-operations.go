package base

import (
	"encoding/json"
	"fmt"
	"log"
)

// User represents a user with basic information
// omitempty tag is used to skip empty fields when encoding to JSON
type User struct {
	ID        int                    `json:"id"`
	Name      string                 `json:"name"`
	Email     string                 `json:"email"`
	Roles     []string               `json:"roles,omitempty"`
	IsActive  bool                   `json:"is_active"`
	Settings  map[string]interface{} `json:"settings,omitempty"`
	CreatedAt string                 `json:"created_at,omitempty"`
}

// JSONOperations demonstrates JSON encoding and decoding
func JSONOperations() {
	// Creating a sample user
	user := User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john@example.com",
		Roles:    []string{"admin", "user"},
		IsActive: true,
		Settings: map[string]interface{}{
			"theme":         "dark",
			"notifications": true,
			"language":      "en",
		},
	}

	// Encoding (Marshal): Converting struct to JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}
	fmt.Printf("Marshaled JSON:\n%s\n\n", jsonData)

	// Pretty print JSON (indented)
	prettyJSON, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatalf("Error creating pretty JSON: %v", err)
	}
	fmt.Printf("Pretty JSON:\n%s\n\n", prettyJSON)

	// Decoding (Unmarshal): Converting JSON to struct
	jsonStr := `{
		"id": 2,
		"name": "Jane Smith",
		"email": "jane@example.com",
		"roles": ["user"],
		"is_active": true,
		"settings": {
			"theme": "light",
			"notifications": false
		},
		"created_at": "2024-02-07T10:30:00Z"
	}`

	var newUser User
	if err := json.Unmarshal([]byte(jsonStr), &newUser); err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}
	fmt.Printf("Unmarshaled struct:\n%+v\n\n", newUser)

	// Working with dynamic JSON (map[string]interface{})
	var dynamicData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &dynamicData); err != nil {
		log.Fatalf("Error unmarshaling to map: %v", err)
	}
	fmt.Printf("Dynamic data (map):\nName: %v\nEmail: %v\n\n",
		dynamicData["name"],
		dynamicData["email"])

	// Handling partial JSON updates
	partialJSON := []byte(`{"name": "Jane Doe", "settings": {"theme": "system"}}`)
	var updatedUser User
	if err := json.Unmarshal(partialJSON, &updatedUser); err != nil {
		log.Fatalf("Error with partial update: %v", err)
	}
	fmt.Printf("Partial update result:\n%+v\n", updatedUser)
}
