package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// go run cmd/exp/bcrypt/bcrypt.go hash "some password here"
// go run cmd/exp/bcrypt/bcrypt.go compare "some password here" "some hash here"

// !!! зашрузили
// go get -u golang.org/x/crypto/bcrypt

func main() {
	switch os.Args[1] {
	case "hash":
		hash(os.Args[2])
	case "compare":
		compare(os.Args[2], os.Args[3])
	default:
		fmt.Printf("Invalid command %v\n", os.Args[1])
	}

}

func hash(password string) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("error hashing: %v\n", err)
		return
	}
	fmt.Println(string(hashedBytes))
}

func compare(password, hash string) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil{
		fmt.Println("Password is invalis", err)
	}
	fmt.Println("Password is correct")

}
