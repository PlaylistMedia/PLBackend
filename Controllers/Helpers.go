package backend

import (
	"crypto/sha512"
	"encoding/hex"
	"time"

	"github.com/golang-jwt/jwt"
)


type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Expire in a month
var expiration_time = time.Now().Add(720 * time.Hour)

func GenerateJWT(user *User) (string, error) {
	// JWT Key is the user's password
	jwt_key := user.Password
	// Create the JWT claims, which includes the username and expiry time
	claims := new(Claims)
	claims.Username= user.Username
	claims.StandardClaims = jwt.StandardClaims{
		// In JWT, the expiry time is expressed as unix milliseconds
		ExpiresAt: expiration_time.Unix(),
	};
	
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		
	// Create the JWT string
	token_string, err := token.SignedString(jwt_key)

	return token_string, err;
}

func HashPassword(user *User) (string, error){
	hash := sha512.New()
	_, err := hash.Write([]byte(user.Password))
	
	return hex.EncodeToString(hash.Sum(nil)), err
}