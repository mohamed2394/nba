package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

const port = "8080"

type UserHandler struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (u *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	randomID := rand.Intn(200) + 1

	resp, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", randomID))
	if err != nil {
		http.Error(w, "Failed to fetch user data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(u); err != nil {
		http.Error(w, "Failed to decode user data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if !u.Completed {
		response := map[string]string{"error": "complete your task"}
		data, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(data)
	} else {
		response := map[string]string{"success": "thanks for completing your task"}
		data, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(data)
	}
}

func main() {
	mux := http.NewServeMux()
	userHandler := &UserHandler{}

	mux.Handle("/api/checkuser", userHandler)

	serv := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Starting server on :%s...", port)
	if err := serv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
