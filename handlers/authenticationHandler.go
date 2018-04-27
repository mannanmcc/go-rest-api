package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/mannanmcc/rest-api/models"
)

//JwtToken toke struct
type JwtToken struct {
	Token string `json:"token"`
}

var signKey = []byte("secret")

//GetToken generate a token for authentication
func (env Env) GetToken(w http.ResponseWriter, r *http.Request) {

	var user models.User
	var tokenString string

	//deserialize and add posted data to User struct
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Error in request:%v", err)
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

/*
ValidateTokenMiddleware validates the token passed to authenticate the user
*/
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
