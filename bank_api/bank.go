package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type account struct {
	Number  string `json:"AccountNumber"`
	Balance string `json:"Balance"`
	Desc    string `json:"AccountDescription"`
}

var accounts []account

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to our bank!")
	fmt.Println("Endpoint: /")
}

func returnAllAccounts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(accounts)
	fmt.Println("Endpoint: /accounts")
}

func returnAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["number"]
	for _, account := range accounts {
		if account.Number == key {
			json.NewEncoder(w).Encode(account)
		}
	}
	fmt.Println("Endpoint: /account/{number}")
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var acc account
	json.Unmarshal(reqBody, &acc)
	accounts = append(accounts, acc)
	json.NewEncoder(w).Encode(acc)
}

func deleteAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["number"]
	for index, acc := range accounts {
		if acc.Number == id {
			accounts = append(accounts[:index], accounts[index+1:]...)
		}
	}
}

func handleRequests() {
	// create a router to handle our requests from the mux package
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/accounts", returnAllAccounts)
	router.HandleFunc("/account/{number}", returnAccount)
	router.HandleFunc("/account", createAccount).Methods("POST")
	router.HandleFunc("/account/{number}", deleteAccount).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	accounts = []account{
		account{Number: "C45t34534", Balance: "24545.5", Desc: "Checking Account"},
		account{Number: "S3r53455345", Balance: "444.4", Desc: "Saving Account"},
	}
	handleRequests()
}
