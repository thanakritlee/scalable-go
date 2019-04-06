package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/thanakritlee/scalable-go/entryserver/mq"
	u "github.com/thanakritlee/scalable-go/entryserver/utils"
)

// response is the structure type template for all response.
type response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

// request for the api request call.
type request struct {
	Route     string `json:"route"`
	Operation string `json:"operation"`
	Body      user   `json:"body"`
}

// user information object.
type user struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Firstname  string `json:"firstname"`
	Middlename string `json:"middlename"`
	Lastname   string `json:"lastname"`
}

// CreateUser is a controller to create and store a User.
func CreateUser(w http.ResponseWriter, r *http.Request) {

	user := user{}

	defer r.Body.Close()

	decorder := json.NewDecoder(r.Body)
	err := decorder.Decode(&user)
	if isError := u.CheckError(err, w); isError {
		return
	}

	req := request{
		Route:     "/api/users",
		Operation: "post",
		Body:      user,
	}

	reqByte, err := json.Marshal(req)
	if isError := u.CheckError(err, w); isError {
		return
	}

	resp, err := mq.Emit(reqByte)
	if isError := u.CheckError(err, w); isError {
		return
	}

	// Unmarshal the payload peer response from chaincode.
	response := response{}
	err = json.Unmarshal(resp, &response)
	if isError := u.CheckError(err, w); isError {
		return
	}

	// Return HTTP response.
	res := u.Message("success")
	res["data"] = response.Data
	res["status"] = response.Status
	u.Response(w, res)

}
