package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/skygeario/skygear-server/uuid"
)

func main() {
	startTime := time.Now()
	claims := jwt.StandardClaims{
		Id:       uuid.New(),
		IssuedAt: time.Now().Unix(),
		Issuer:   "myappname",
		Subject:  "myuserid",
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("secret"))
	fmt.Println(tokenString, err)
	fmt.Println(time.Now().Sub(startTime))

}
