package resource

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"golnfuturecapacities/api/middleware"
	"golnfuturecapacities/api/models"
	"golnfuturecapacities/api/service"
	"golnfuturecapacities/api/utils"
	"net/http"
	"strconv"
	"time"
)

type UserResource struct {
	UserService service.UserService
}

func UserController(userService service.UserService) *UserResource {
	return &UserResource{
		UserService: userService,
	}
}
func (app *UserResource) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var payload models.User
	var userRoles models.UserRole
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	// validate payload
	if err := utils.Validate.Struct(&payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error", errors))
		return
	}
	// validate regex
	_, err := utils.RegexValidate(payload.FirstName)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Field payload %s", err))
		return
	}
	// validate regex
	_, err = utils.RegexValidate(payload.LastName)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Field payload payload %s", err))
		return
	}
	// check if user email exist
	user, _ := app.UserService.Exists(payload.Email)
	if payload.Email == user.Email {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with that email %s already exists", payload.Email))
		return
	}
	// save user payload
	save, err := app.UserService.Save(&payload)
	if err != nil {
		utils.WriteError(w, http.StatusUnprocessableEntity, fmt.Errorf("invalid payload %s", err))
		return
	}
	// add to user role
	userRoles.UserId = uint(save.ID)
	_, err = app.UserService.AddToUserRole(&userRoles)
	if err != nil {
		utils.WriteError(w, http.StatusUnprocessableEntity, fmt.Errorf("invalid payload entry %s", err))
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"data": save,
	})
}
func (app *UserResource) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var payload models.Login
	var bodyPayload models.User
	//var vtoken models.JWToken

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	// validate payload
	if err := utils.Validate.Struct(&payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error", errors))
		return
	}
	// check if user email exist
	user, _ := app.UserService.Exists(payload.Email)
	if user.Email == bodyPayload.Email {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid username/password"))
		return
	}
	// check if user is enabled
	enabled, _ := app.UserService.IsEnabled(payload.Email)
	if enabled.Enabled.Bool == false {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user not enabled"))
		return
	}
	// check hash password
	password := bodyPayload.Password
	hashPwd := user.Password
	compare := utils.ComparePassword(password, hashPwd)
	token, _ := middleware.GenerateAccessToken(user)
	refresh, _ := middleware.GenerateRefreshToken(user)

	if compare {
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   "",
			Expires: time.Now().Add(24 * time.Hour),
		})
	}
	response := map[string]any{
		"user":          user,
		"access_token":  token,
		"refresh_token": refresh,
	}
	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"data": response,
	})
}
func (app *UserResource) FindUserIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	params := r.PathValue("id")
	userId, err := strconv.ParseUint(params, 10, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("that role id not avaliable"))
		return
	}
	user, err := app.UserService.Find(userId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("internel server error"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"data": user,
	})
}
func (app *UserResource) FindAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	user, err := app.UserService.FindAll()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("internel server error"))
		return
	}
	//utils.WriteJSON(w, http.StatusOK, map[string]interface{}{"data": user})
	utils.WriteJSON(w, http.StatusOK, user)
}
func (app *UserResource) UpdateUserIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	params := r.PathValue("id")
	userId, err := strconv.ParseUint(params, 10, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("that role id not avaliable"))
		return
	}
	//var payload models.Role
	userPayload := &models.User{}
	if err := utils.ParseJSON(r, &userPayload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	// validate regex
	_, err = utils.RegexValidate(userPayload.LastName)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Field payload %s", err))
		return
	}
	// set Role.Id = request = roleId
	userPayload.ID = userId
	err = app.UserService.Update(userPayload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("internel server error"))
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"success": "Update user successfully.",
	})
}
func (app *UserResource) DeleteUserIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	params := r.PathValue("id")
	userId, err := strconv.ParseUint(params, 10, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("that role id not avaliable"))
		return
	}
	err = app.UserService.Delete(userId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("that role id not avaliable"))
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"success": "Delete user successfully.",
	})
}
