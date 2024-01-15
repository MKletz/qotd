package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AllQuotes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(quotes)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "qotd Oct 11 2023 10:41 am")
}

func OneQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["quoteId"])
	if err == nil {
		json.NewEncoder(w).Encode(quotes[id])
	}
}

func RandomQuote(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(quotes[rand.Intn(7)])
}

func Version(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "3.0.0")
}

func WrittenIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GO")
}
