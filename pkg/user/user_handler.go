package user

import (
	"net/http"
	"handsongo/pkg/utils"
	"log"
	"encoding/json"
)

type UserHandler struct {
	userService Service
}

func NewUserHandler(userService Service) *UserHandler {
	return &UserHandler{
		userService : userService,
	}
}


func (uh *UserHandler) CreateUser(resp http.ResponseWriter, req *http.Request) {
	var ur UserRequest
	err := utils.ParseRequest(req, &ur)

	if err != nil {
		resp.WriteHeader(500)
		resp.Write([]byte("request parse error"))
		return
	}
	
	user,_:= NewUser(ur)
	log.Printf("ur.FullName=====>"+user.FullName)
	res, _ := uh.userService.CreateUser(req.Context(),&user);
	userResp := UserResp{
		UserID  : res.ID,
		Message : "user is successfully created",
	}
	result, _ := json.Marshal(&userResp)
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(result)
}