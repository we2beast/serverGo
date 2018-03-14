package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
	"Calday-Server/handlers"
)

var secretKey = []byte(os.Getenv("SESSION_SECRET"))
var usersOne = map[string]string{"naren": "passme", "admin": "password"}

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	tokenString, err := request.HeaderExtractor{"access_token"}.ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Access Denied; Please check the access token"))
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		response := make(map[string]interface{})
		response["time"] = time.Now().String()
		response["id"] = claims["id"]
		responseJSON, _ := json.Marshal(response)
		w.Write(responseJSON)
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
	}
}

func middlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := request.HeaderExtractor{"access_token"}.ExtractToken(r)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return secretKey, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Access Denied; Please check the access token"))
			return
		} else {
			next.ServeHTTP(w, r)
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			response := make(map[string]interface{})
			response["time"] = time.Now().String()
			response["time"] = time.Now().String()
			response["id"] = claims["id"]
			responseJSON, _ := json.Marshal(response)
			w.Write(responseJSON)
		} else {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(err.Error()))
		}
	})
}

func main() {
	r := mux.NewRouter()
	// r.HandleFunc("/getToken", getTokenHandler)

	// Auth
	r.HandleFunc("/api/register", handlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/api/login", handlers.LoginHandler).Methods("POST")

	r.HandleFunc("/healthcheck", HealthcheckHandler)
	r.HandleFunc("/api/index", handlers.IndexHandler).Methods("GET")

	// User
	r.Handle("/api/user/{id}", middlewareAuth(http.HandlerFunc(handlers.UserHandler))).Methods("GET")
	r.HandleFunc("/api/user/edit", handlers.UserEditHandler).Methods("POST")

	// Events
	r.HandleFunc("/api/events/{id}", handlers.EventHandler).Methods("GET")
	r.HandleFunc("/api/events/insert", handlers.EventInsertHandler).Methods("POST")
	r.HandleFunc("/api/events/update", handlers.EventUpdateHandler).Methods("POST")

	http.Handle("/", r)
	srv := &http.Server{
		Handler: r,
		Addr: "127.0.0.1:8001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
