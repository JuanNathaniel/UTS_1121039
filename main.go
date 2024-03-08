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

	//Get all rooms
	router.HandleFunc("/getrooms", controllers.GetAllRooms).Methods("GET")

	//Get detail rooms
	router.HandleFunc("/getrooms/getdetail", controllers.GetDetailRooms).Methods("GET")

	http.Handle("/", router)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
