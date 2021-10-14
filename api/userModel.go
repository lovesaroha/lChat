/*  Love Saroha
    lovesaroha1994@gmail.com (email address)
    https://www.lovesaroha.com (website)
    https://github.com/lovesaroha  (github)
*/
package api

import (
	"encoding/json"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// User object.
type userObject struct {
	EmailAddress string `json:"emailAddress,omitempty" bson:"emailAddress,omitempty"`
	Token        string `json:"token,omitempty" bson:"token,omitempty"`
}

// User token structure defined here.
type userTokenObject struct {
	EmailAddress string `json:"emailAddress"`
	jwt.StandardClaims
}

// This function return json string of user.
func userJSON(user userObject) string {
	user.Token = encryptUserToken(user.EmailAddress)
	s, _ := json.Marshal(user)
	return string(s)
}

// This function encrypts user's email address using key.
func encryptUserToken(emailAddress string) string {
	expirationTime := time.Now().Add(12 * time.Hour)
	claims := &userTokenObject{
		EmailAddress: emailAddress,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("qwertyuiopasdfghjklzxcvbnmlkjhgf"))
	return tokenString
}

// This function decrypt user's email address using key.
func invalidUserToken(user *userObject) bool {
	claims := &userTokenObject{}
	if token, err := jwt.ParseWithClaims(user.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("qwertyuiopasdfghjklzxcvbnmlkjhgf"), nil
	}); err != nil || !token.Valid {
		return true
	}
	user.EmailAddress = claims.EmailAddress
	return false
}
