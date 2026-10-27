package controller

import (
	"core-banking/internal/core/user/dto"
	"core-banking/internal/core/user/service"
	"core-banking/middleware"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type UserControllerImp struct {
	authService service.AuthService
	userService service.UserService
}

func NewUserController(
	authsService service.AuthService,
	userService service.UserService,
) UserController {
	return &UserControllerImp{
		authService: authsService,
		userService: userService,
	}
}

func (controller *UserControllerImp) Register(writer http.ResponseWriter, request *http.Request) {

	ctx := request.Context()

	var req dto.RegisterRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		writer.Header().Set("content-type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(dto.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   "Input Salah",
		})
		return
	}
	user, err := controller.authService.Register(ctx, req)
	if err != nil {
		status := http.StatusInternalServerError
		msg := "internal server error"

		switch {
		case errors.Is(err, service.ErrInvalidInput):
			status = http.StatusBadRequest
			msg = "Username & password are required"
		case errors.Is(err, service.ErrUsernameExists):
			status = http.StatusConflict
			msg = "username already exists"
		}
		writer.WriteHeader(status)
		json.NewEncoder(writer).Encode(dto.Response{
			Code:   status,
			Status: http.StatusText(status),
			Data:   msg,
		})
		return
	}
	writer.Header().Set("content-type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(dto.Response{
		Code:   http.StatusCreated,
		Status: "create",
		Data:   user,
	})
}

func (controller *UserControllerImp) Login(writer http.ResponseWriter, request *http.Request) {

	ctx := request.Context()

	var input dto.LoginRequest

	if err := json.NewDecoder(request.Body).Decode(&input); err != nil {
		writer.Header().Set("content-type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(dto.Response{
			Code:   http.StatusBadRequest,
			Status: "400",
			Data:   "Data salah",
		})
		return
	}
	user, err := controller.authService.Login(ctx, input)
	if err != nil {
		writer.Header().Set("content-type", "application/json")
		if err == service.ErrInvalidCredential {
			writer.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(writer).Encode(dto.Response{
				Code:   http.StatusUnauthorized,
				Status: "401",
				Data:   "Failed",
			})
			return
		}
		writer.Header().Set("content-type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(dto.Response{
			Code:   http.StatusInternalServerError,
			Status: "500",
			Data:   "server error",
		})
		return
	}
	writer.Header().Set("content-type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(dto.Response{
		Code:   http.StatusCreated,
		Status: "201",
		Data:   user,
	})
}

func (controller *UserControllerImp) Profile(writer http.ResponseWriter, request *http.Request) {

	ctx := request.Context()

	userID, ok := request.Context().Value(middleware.UserIDKey).(int64)
	if !ok {
		writer.Header().Set("content-type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(writer).Encode(dto.Response{
			Code:   http.StatusUnauthorized,
			Status: "401",
			Data:   "failed",
		})
		return
	}
	_, user, err := controller.userService.Profile(ctx, userID)
	if err != nil {
		fmt.Println("profile error", err)
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}
	writer.Header().Set("content-type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(dto.Response{
		Code:   http.StatusCreated,
		Status: "201",
		Data:   user,
	})
}
