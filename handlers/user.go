package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/roaris/go-sns-api/httputils"

	"github.com/go-sql-driver/mysql"
	"github.com/roaris/go-sns-api/models"
	"gopkg.in/go-playground/validator.v9"
)

type UserRequest struct {
	Name     string
	Email    string
	Password string
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	// application/jsonのみ受け付ける
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// リクエストボディをUserRequestに変換する
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var userRequest UserRequest
	json.Unmarshal(body, &userRequest)

	user, err := models.CreateUser(userRequest.Name, userRequest.Email, userRequest.Password)
	if _, ok := err.(validator.ValidationErrors); ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if _, ok := err.(*mysql.MySQLError); ok {
		w.WriteHeader(http.StatusConflict)
		return
	} else if err != nil && err.Error() == "too short password" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(user.SwaggerModel())
	w.Write(res)
}

func GetLoginUser(w http.ResponseWriter, r *http.Request) {
	userID := httputils.GetUserIDFromContext(r.Context())
	user, _ := models.GetUserById(userID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(user.SwaggerModelWithEmail())
	w.Write(res)
}
