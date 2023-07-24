package middleware

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type TokenClaims struct {
	UserId uint
	RoleId []int
	jwt.StandardClaims
}

const TokenSecret = "my-secret-key"

func VerifyToken(tokenString, origin string) (bool, *TokenClaims) {
	claims := &TokenClaims{}
	token, _ := getTokenFromString(tokenString, claims, origin)
	if token.Valid {
		if e := claims.Valid(); e == nil {
			return true, claims
		}
	}
	return false, claims
}

func getTokenFromString(tokenString string, claims *TokenClaims, origin string) (*jwt.Token, error) {
	// token, err := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
	// 	// Validate the alg is what you expect:
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	// 	}
	// 	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	// 	return []byte(secretKey), nil
	// })

	// if err != nil {
	// 	return nil, err
	// }

	// if !token.Valid {
	// 	return nil, errors.New("Invalid token")
	// }

	// return token, nil
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(origin), nil
	})
}

func GenerateToken(userId uint, roleId []int) (string, error) {
	// Set token claims
	claims := &TokenClaims{
		UserId: userId,
		RoleId: roleId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token will expire in 24 hours
		},
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(TokenSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
