package models

type Accounts struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type Games struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Max_player int    `json:"max_player"`
}

type Rooms struct {
	ID        int    `json:"id"`
	Room_name string `json:"room_name"`
	Id_game   string `json:"id_game"`
}

type Participant struct {
	ID         int `json:"id"`
	Id_room    int `json:"id_room"`
	Id_account int `json:"id_account"`
}

type RoomResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Rooms  `jsonL:"data"`
}
type RoomsResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Rooms `json:"data"`
}
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
