package controllers

import (
	m "UTS_1121039/models"
	"encoding/json"
	"log"
	"net/http"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM rooms"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		//send error response
		return
	}

	var room m.Rooms
	var rooms []m.Rooms
	var response m.RoomsResponse
	var error m.ErrorResponse

	for rows.Next() {
		if err := rows.Scan(&room.ID, &room.Room_name, &room.Id_game); err != nil {
			log.Println(err)
			//send error response
			error.Status = 400
			error.Message = "Get rooms Failed"
			error.Message = err.Error()
			json.NewEncoder(w).Encode(error)
			return
		} else {
			rooms = append(rooms, room)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	// var response UsersResponse
	// if len(users) < 5

	response.Status = 200
	response.Message = "Success"
	response.Data = rooms
	json.NewEncoder(w).Encode(response)

}

func GetDetailRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Parse input dari request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}

	var query string

	roomID := r.Form.Get("id")

	if roomID == " " {
		http.Error(w, "please provide ROOM ID", http.StatusBadRequest)
		return
	}

	query = "SELECT r.id AS room_id, r.room_name, g.id AS game_id, g.name AS game_name, p.id AS participant_id FROM rooms r JOIN games g ON r.id_game = g.id JOIN participants p ON r.id = p.id_room"

	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Failed to execute query", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var room m.Rooms
	var game m.Games
	var participant m.Participant

	var rooms []m.Rooms
	var response m.RoomsResponse
	for rows.Next() {

		err := rows.Scan(&room.ID, &room.Room_name, &game.ID, &game.Name, &participant.ID)
		if err != nil {
			http.Error(w, "Failed to scan row", http.StatusInternalServerError)
			return
		} else {
			rooms = append(rooms, room)
			response.Status = 200
			response.Message = "Success"
			response.Data = rooms
			json.NewEncoder(w).Encode(rooms)
		}

	}

	w.Header().Set("Content-Type", "application/json")
	response.Status = 200
	response.Message = "Success"
	response.Data = rooms
	json.NewEncoder(w).Encode(response)
}

func InsertRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// Parse input dari request
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}

	// Username := r.Form.Get("username")

	// if err != nil {
	// 	http.Error(w, "failed to parse form data", http.StatusBadRequest)
	// 	return
	// }
	// username := r.Form.Get("username")
	// if username == "" {
	// 	http.Error(w, "please provide both email and password", http.StatusBadRequest)
	// 	return
	// }

	// accounts := connectgorm()

	// user := models.Accounts{Username: username}

	// result := db.Create(&accounts)
	// if result.Error != nil {
	// 	http.Error(w, result.Error.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// var response = m.UserResponse{}
	// response.Status = http.StatusOK
	// response.Message = "Success"
	// response.Data = user
	// json.NewEncoder(w).Encode(response)
}
