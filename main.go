package main

import (
	"fmt"

	jwtvalidator "github.com/DakshithaS/test-policy-hub/jwt-validator/v2.1.0"
)

func main() {
	validator := jwtvalidator.New("secret")
	fmt.Println("JWT validator created")
	valid, err := validator.ValidateWithContext(nil, "some.token.here")
	if err != nil {
		fmt.Println("Error:", err)
	} else if valid {
		fmt.Println("Token is valid")
	} else {
		fmt.Println("Token is invalid")
	}
}