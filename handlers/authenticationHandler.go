package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/mannanmcc/rest-api/models"
)

//JwtToken toke struct
type JwtToken struct {
	Token string `json:"token"`
}

var signKey = []byte("secret")

//GetToken - generate a token for authentication
func (env Env) GetToken(w http.ResponseWriter, r *http.Request) {

	var user models.User
	var tokenString string

	//deserialize and add posted data to User struct
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Oops, Error in the request", http.StatusForbidden)
		return
	}

	//todo - validate user details with database
	userRepo := models.UserRepository{Db: env.Db}
	_, err = userRepo.FindByUserNameAndPassword(user.Username, user.Password)
	if err != nil {
		log.Println("Oops, wrong credential passed!")
		http.Error(w, "Oops, wrong credential passed!", http.StatusForbidden)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	})

	tokenString, err = token.SignedString(signKey)
	if err != nil {
		log.Println("Oops, there is a error to create your token")
		http.Error(w, "there was an error in login", http.StatusInternalServerError)
		return
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
			log.Println("Token is not valid")
			http.Error(w, "Oops, wrong credential passed!", http.StatusUnauthorized)
		}
	} else {
		log.Println(w, "Unauthorized access to this resource")
		http.Error(w, "Oops, wrong credential passed!", http.StatusUnauthorized)
	}
}
