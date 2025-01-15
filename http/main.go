package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json: "id"`
	Name string `json: "name,omitempty"`
}

var (
	users = []User{{1, "Vasya"}, {2, "Petya"}}
)

func main() {
	http.HandleFunc("/users", authMiddleware(loggerMiddleware(hanleUsers)))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("x-id")
		if userID == "" {
			log.Printf("[%s] %s - error: userID is not provided\n", r.Method, r.RequestURI)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "userID", userID)

		r = r.WithContext(ctx)
		next(w, r)
	}
}

func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ifFromContext := r.Context().Value("userID")
		userID, ok := ifFromContext.(string)

		if !ok {
			log.Printf("[%s] %s - error: userID is invalid\n", r.Method, r.URL)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Printf("[%s] %s - userID: %s\n", r.Method, r.URL, userID)

		log.Printf("Request: %s %s\n", r.Method, r.URL)
		next(w, r)
	}
}

func hanleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		addUser(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(resp)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user User
	if err := json.Unmarshal(reqBytes, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	users = append(users, user)
	w.WriteHeader(http.StatusCreated)
}
