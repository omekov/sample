package customers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/omekov/sample/internal/apiserver/models"
	"golang.org/x/crypto/bcrypt"
)

func createJWT(customer *models.Customer, auth *models.SignInput) (string, error) {
	err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(auth.Password))
	if err != nil {
		return "", err
	}
	expirationTime := time.Now().Add(5 * time.Minute).Unix()
	customer.Password = ""
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.Claims{
		Customer: *customer,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	})
	return token.SignedString([]byte("secret"))
}
func encryptString(p string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// func verifyToken(tokenString string) (string, error) {
// 	claims := jwt.MapClaims{}
// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		return jwtSecretKey, nil
// 	})
// 	if err != nil {
// 		return "", err
// 	}
// 	return token, nil
// }
