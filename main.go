package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"main.go/connect"
	"main.go/structures"
)

func main() {
	connect.InitializeDB()
	defer connect.CloseConnection()

	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", GetUser).Methods("GET")
	r.HandleFunc("/user/new", NewUser).Methods("POST")
	r.HandleFunc("/user/update/{id}", UpdateUser).Methods("PATCH")
	r.HandleFunc("/user/delete/{id}", DeleteUser).Methods("DELETE")

	log.Println("The server is running on port 8098")
	log.Fatal(http.ListenAndServe(":8098", r))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	status := "success"
	var message string
	user := connect.GetUser(userId)

	if user.Id <= 0 {
		status = "error"
		message = "User not found"
	}

	response := structures.Response{
		Status:  status,
		Data:    user,
		Message: message,
	}

	json.NewEncoder(w).Encode(response)
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	user := GetUserRequest(r)
	connect.CreateUser(user)

	response := structures.Response{
		Status:  "Success",
		Data:    connect.CreateUser(user),
		Message: "User added",
	}

	json.NewEncoder(w).Encode(response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	user := GetUserRequest(r)

	response := structures.Response{
		Status:  "Success",
		Data:    connect.UpdateUser(userId, user),
		Message: "User updated",
	}

	json.NewEncoder(w).Encode(response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	var user structures.User

	connect.DeleteUser(userId)

	response := structures.Response{
		Status:  "Success",
		Data:    user,
		Message: "User deleted",
	}

	json.NewEncoder(w).Encode(response)
}

func GetUserRequest(r *http.Request) structures.User {
	var user structures.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	return user
}
