package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name        string
	Description string
	Age         int
	Friends     []User
}

func main() {

	userList := `[
		{
		  "name": "Sparsh",
		  "desciption": "Software Developer",
		  "age": 22,
		  "friends": [
			{
				"name":"Amber",
				"description":"Frontend Developer",
				"age":22,
				"friends":[]
			}
		  ]
		},
		{
		  "name":"Amber",
		  "description":"Frontend Developer",
		  "age":22,
		  "friends":[
			{
				"name": "Sparsh",
				"description": "Software Developer",
				"age": 22,
				"friends":[]
			}
		  ]
		}
	  ]`

	var users []User

	json.Unmarshal([]byte(userList), &users)

	myRouter := mux.NewRouter().StrictSlash(true)

	//get all Users
	getUsersDetails := func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(users)
	}
	// get User by Name
	getUserName := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["Name"]

		for _, user := range users {
			if user.Name == name {
				json.NewEncoder(w).Encode(user)
				fmt.Println("Matched")
			}
		}
	}
	// remove User by Name
	removeUser := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["Name"]

		for index, user := range users {
			if user.Name == name {
				users = append(users[:index], users[index+1:]...)
				json.NewEncoder(w).Encode(users)
				fmt.Println("Deleted")
			}
		}
	}

	myRouter.HandleFunc("/users", getUsersDetails).Methods("GET")

	myRouter.HandleFunc("/users/{Name}", getUserName)

	myRouter.HandleFunc("/users/{Name}", removeUser).Methods("DELETE")

	fmt.Println("Server is running on Port:5000")
	http.ListenAndServe(":5000", myRouter)

}
