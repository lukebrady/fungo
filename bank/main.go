package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Account stores all of the accounts that are created.
type Account struct {
	accounts map[int]string
}

type open struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Mail  string `json:"mail"`
}

// OpenAccount allows a user to open a bank account
func OpenAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("You can open an account with JSON.\n"))
		w.Write([]byte("Send with name, phone number, and email.\n"))
	} else if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		data := new(open)
		json.Unmarshal(body, data)

		if data.Name != "" && data.Phone != "" && data.Mail != "" {
			fmt.Printf("Hello %s, thank you for opening a bank account with us.\n", data.Name)
			w.Write([]byte("Thank you for opening a bank account, " + data.Name + "!\n"))
		}
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/open", OpenAccount)
	http.ListenAndServe(":8080", mux)
}
