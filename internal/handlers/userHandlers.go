package handlers

import (
	"context"
	"time"

	"pet_project_1_etap/internal/userService"
	"pet_project_1_etap/internal/web/users"
)

type UserHandler struct {
	UserService *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		UserService: service,
	}
}

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.UserService.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}
	for _, u := range allUsers {
		var deletedAt *time.Time
		if u.DeletedAt.Valid {
			deletedAt = &u.DeletedAt.Time
		}

		user := users.User{
			Id:        &u.ID,
			Email:     &u.Email,
			Password:  &u.Password,
			CreatedAt: &u.CreatedAt,
			UpdatedAt: &u.UpdatedAt,
			DeletedAt: deletedAt, // ✅ теперь тип *time.Time
		}
		response = append(response, user)
	}

	return response, nil
}

func (h *UserHandler) PostUsers(_ context.Context, req users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	body := req.Body

	newUser := userService.User{
		Email:    *body.Email,
		Password: *body.Password,
	}

	createdUser, err := h.UserService.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	var deletedAt *time.Time
	if createdUser.DeletedAt.Valid {
		deletedAt = &createdUser.DeletedAt.Time
	}

	response := users.PostUsers201JSONResponse{
		Id:        &createdUser.ID,
		Email:     &createdUser.Email,
		Password:  &createdUser.Password,
		CreatedAt: &createdUser.CreatedAt,
		UpdatedAt: &createdUser.UpdatedAt,
		DeletedAt: deletedAt,
	}

	return response, nil
}

func (h *UserHandler) PatchUsersId(_ context.Context, req users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	id := req.Id
	body := req.Body

	updatedUser := userService.User{
		Email:    *body.Email,
		Password: *body.Password,
	}

	updated, err := h.UserService.UpdateUserByID(uint(id), updatedUser)
	if err != nil {
		return nil, err
	}

	var deletedAt *time.Time
	if updated.DeletedAt.Valid {
		deletedAt = &updated.DeletedAt.Time
	}

	response := users.PatchUsersId200JSONResponse{
		Id:        &updated.ID,
		Email:     &updated.Email,
		Password:  &updated.Password,
		CreatedAt: &updated.CreatedAt,
		UpdatedAt: &updated.UpdatedAt,
		DeletedAt: deletedAt,
	}

	return response, nil
}

func (h *UserHandler) DeleteUsersId(_ context.Context, req users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := req.Id

	err := h.UserService.DeleteUserByID(uint(id))
	if err != nil {
		return nil, err
	}

	return users.DeleteUsersId204Response{}, nil
}
