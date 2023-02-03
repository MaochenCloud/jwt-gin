package token

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// store in the env value
var mySigningKey = []byte("Intel ESPD")

func GenerateToken() (string, error) {

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

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return nil
	} else {
		return fmt.Errorf("Unexpected Claims: %v", token.Claims)
	}
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// func ExtractTokenID(c *gin.Context) (uint, error) {

// 	tokenString := ExtractToken(c)
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(os.Getenv("API_SECRET")), nil
// 	})
// 	if err != nil {
// 		return 0, err
// 	}
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if ok && token.Valid {
// 		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
// 		if err != nil {
// 			return 0, err
// 		}
// 		return uint(uid), nil
// 	}
// 	return 0, nil
// }
