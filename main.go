package main

import (
	"UTS_1121039/controllers"
	"fmt"
	"log"
	"net/http"

	//GORM

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	//"github.com/joho/godotenv"
)

func main() {
	//loadEnv()
	router := mux.NewRouter()

	//GORM
	router.HandleFunc("/v2/insertuser2", controllers.InsertUser2).Methods("POST")
	// //updateuser2
	// router.HandleFunc("/v2/updateuser2", controllers.UpdateUser2).Methods("PUT")
	// //deleteuser2
	// router.HandleFunc("/v2/deleteuser2", controllers.DeleteUser2).Methods("DELETE")
	// //selectuser2
	// router.HandleFunc("/v2/selectuser2", controllers.SelectUser2).Methods("GET")
	// //raw/complex query
	// router.HandleFunc("/v2/updateraw2", controllers.UpdateRaw2).Methods("PUT")

	// //Get all users
	// router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")

	// //Insert new users
	// router.HandleFunc("/insertuser", func(w http.ResponseWriter, r *http.Request) {
	// 	controllers.InsertUser(w, r, "EP", 19, "Bandung")
	// }).Methods("POST")

	// //delete user
	// router.HandleFunc("/deleteuser", func(w http.ResponseWriter, r *http.Request) {
	// 	controllers.DeleteUser(w, r, "feawf", 19)
	// }).Methods("DELETE")

	// //update user
	// router.HandleFunc("/updateuser", func(w http.ResponseWriter, r *http.Request) {
	// 	controllers.UpdateUser(w, r, "haloww", "fefefe")
	// }).Methods("PUT")

	// //Get detail user transaction
	// router.HandleFunc("/detailusertransaction", func(w http.ResponseWriter, r *http.Request) {
	// 	controllers.GetDetailUserTransaction(w, r, 0)
	// }).Methods("GET")

	// //Login
	// router.HandleFunc("/loginuser", controllers.LoginUser).Methods("POST")

	http.Handle("/", router)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
