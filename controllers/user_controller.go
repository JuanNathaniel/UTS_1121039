package controllers

import (
	"UTS_1121039/models"
	m "UTS_1121039/models"
	"encoding/json"
	"net/http"
	"strconv"
)

// GORM Insert
func InsertUser2(w http.ResponseWriter, r *http.Request) {
	// Parse input dari request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}
	name := r.Form.Get("name")
	age, _ := strconv.Atoi(r.Form.Get("age"))
	address := r.Form.Get("address")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	if email == "" || password == "" {
		http.Error(w, "please provide both email and password", http.StatusBadRequest)
		return
	}

	db := connectgorm()

	user := models.User{Name: name, Age: age, Address: address, Email: email, Password: password}

	result := db.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var response = m.UserResponse{}
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = user
	json.NewEncoder(w).Encode(response)
}
