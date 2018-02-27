package handlers

import (
	"strings"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"github.com/mannanmcc/rest-api/models"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go/request"
)

type JwtToken struct {
	Token string `json:"token"`
}

var signKey = []byte("secret")

func (env Env) GetToken(w http.ResponseWriter, r *http.Request) {

	var user models.User
	var tokenString string

	//deserialize and add posted data to User struct
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Error in request")
		return
	}

	//todo - validate user details with database
	if strings.ToLower(user.Username) != "someone" || user.Password != "password" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Println("Error logging in")
			fmt.Fprint(w, "Invalid credentials")
			return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	})

	tokenString, err = token.SignedString(signKey)
	if err != nil {
		fmt.Println(tokenString)
	}
	json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}

func (env Env) ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return signKey, nil
		})

	if err == nil {
		if token.Valid {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
	}
}