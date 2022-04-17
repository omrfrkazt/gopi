package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	. "testApi/dataloaders"
	. "testApi/models"
)

func Run() {
	http.HandleFunc("/user", UserHandler)
	http.ListenAndServe(":8080", nil)
}
func UserHandler(w http.ResponseWriter, r *http.Request) {
	users := LoadUsers(true)
	query := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(query)
	response := BaseResponseModel{}
	if users == nil {
		response.Message = "No users found"
		response.StatusCode = http.StatusNotFound
	} else if id > 0 {
		for _, user := range users {
			if user.Id == id {
				response.Message = "Success"
				response.StatusCode = http.StatusOK
				response.Data = user
			}
		}
	} else {
		response.Message = "Success"
		response.StatusCode = http.StatusOK
		response.Data = users
		response.Message = "user count: " + strconv.Itoa(len(users))
	}
	jsonResponse, _ := json.Marshal(response)
	w.WriteHeader(response.StatusCode)
	w.Write([]byte(jsonResponse))
}
