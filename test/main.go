package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySigningKey = []byte("Intel EOMS")

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func BuildToken() (string, error) {

	claims := MyCustomClaims{
		Username: "intel",
		StandardClaims: jwt.StandardClaims{
			Audience:  "somebody_else",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)).Unix(),
			Id:        "1",
			IssuedAt:  jwt.NewNumericDate(time.Now()).Unix(),
			Issuer:    "intel",
			NotBefore: jwt.NewNumericDate(time.Now()).Unix(),
			Subject:   "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return ss, err
}

//verification

func PasreToken(ss string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(ss, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {

		return claims, nil
	} else {
		fmt.Println(err)
		return nil, err
	}
}
func main() {
	ss, err := BuildToken()
	if err != nil {
		return
	}
	fmt.Printf("Token: %v\n", ss)

	claims, err := PasreToken(ss)
	if err != nil {
		return
	}
	fmt.Printf("%v %v\n", claims.Username, claims.StandardClaims.Issuer)
}
