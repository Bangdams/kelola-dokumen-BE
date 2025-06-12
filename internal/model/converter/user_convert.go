package converter

import (
	"log"

	"github.com/Bangdams/kelola-dokumen-BE/internal/entity"
	"github.com/Bangdams/kelola-dokumen-BE/internal/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	log.Println("log from user to response")

	return &model.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Role:     user.Role,
		NameRw:   user.RwList.NameRw,
	}
}

func LoginUserToResponse(token string) *model.LoginResponse {
	log.Println("log from login user to response")

	return &model.LoginResponse{
		AccessToken: token,
	}
}

func UserToResponses(users *[]entity.User) *[]model.UserResponse {
	var userResponses []model.UserResponse

	log.Println("log from user to responses")

	for _, user := range *users {
		userResponses = append(userResponses, *UserToResponse(&user))
	}

	return &userResponses
}
