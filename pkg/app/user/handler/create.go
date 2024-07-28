package user

import (
	"encoding/json"
	"net/http"
	"log"
	"Project/pkg/app/user/model"
	"gopkg.in/validator.v2"
	"Project/pkg/api"
)



func (h UserHandler) CreatUserHandler(w http.ResponseWriter, r *http.Request)  {
	var user model.CreateUserModel

	response := make(map[string]bool)
	
	body := r.Body
	defer body.Close()
	
	err := json.NewDecoder(body).Decode(&user)

	
	
	if err := validator.Validate(user); err != nil {
		api.Error(err, http.StatusBadRequest).Send(w)
		return
	}

	
	if err != nil {	
		log.Fatal(err.Error())
	} 
	succes, err := h.UserService.CreateUserService(user)
	if err != nil {
		api.Error(err, http.StatusInternalServerError).Send(w)
	} 
	
	response["SUCCESS"]= succes

	jsonString, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err.Error())
	} 
	w.Write(jsonString)
}