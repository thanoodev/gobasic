package main

import (
	"gobasic/restapi"
)

func main() {
	restapi.Init()

	// Curl command examples to test the API:
	// Get Profile : curl -X GET http://localhost:8080/profile
	// Add Profile : curl -X POST http://localhost:8080/addprofile -d '{"id": 10, "name": "Nui"}'
}
