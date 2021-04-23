package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	signingString = "testq2eqsd8f76sad8f7a6s dfasudyftr" //todo use value from config
)

func NewToken(userId, userRole string) (string, error) {
	mySigningKey := []byte(signingString)

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + 60*60*24,
		Subject:   userId,
		Audience:  userRole,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	return ss, err
}

func VerifyToken(tokenString string) (*jwt.StandardClaims, error) {
	mySigningKey := []byte(signingString) //todo use value from config

	var claims jwt.StandardClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	//todo work on messages below
	if token.Valid {
		fmt.Println("You look nice today")
		return &claims, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
		return nil, err
	} else {
		fmt.Println("Couldn't handle this token:", err)
		return nil, err
	}

}