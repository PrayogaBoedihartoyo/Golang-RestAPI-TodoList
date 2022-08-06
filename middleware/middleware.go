package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"log"
	"main/helper"
	"net/http"
)

var (
	router    *mux.Router
	secretkey string = "secretkeyjwt"
)

// check whether user is authorized or not
func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Header["Token"] == nil {
			var err helper.Error
			err = helper.SetError(err, "No Token Found")
			json.NewEncoder(w).Encode(err)
			return
		}

		var mySigningKey = []byte(secretkey)

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing token.")
			}
			return mySigningKey, nil
		})

		if err != nil {
			var err helper.Error
			err = helper.SetError(err, "Your Token has been expired.")
			json.NewEncoder(w).Encode(err)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims);
		if ok && token.Valid {
			var reserr helper.Error
			reserr = helper.SetError(reserr, "Not Authorized.")
			json.NewEncoder(w).Encode(err)
			return
		}
		log.Println(claims)
		handler.ServeHTTP(w, r)
		return
	}
}
