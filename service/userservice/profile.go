package userservice

import (
	"fmt"
	"quizofkingz/entity"
)

type ProfileRequest struct {
	UserID uint `json:"user_id"`
}

type ProfileResponse struct {
	User entity.User `json:"user"`
}

func (s Service) Profile(req ProfileRequest) (ProfileResponse, error) {
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return ProfileResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return ProfileResponse{User: user}, nil
}
