package delivery

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ell1jah/db_cp/internal/models"
	"github.com/ell1jah/db_cp/pkg/logger"

	"github.com/asaskevich/govalidator"
)

type TokenForm struct {
	Token string `json:"token"`
}

type loginForm struct {
	Login    string `valid:"minstringlength(5)" json:"login"`
	Password string `valid:"minstringlength(5)" json:"password"`
}

type UserService interface {
	CreateUser(models.User) (int, error)
	GetUserByLoginAndPassword(string, string) (models.User, error)
}

type SessionManager interface {
	CreateSession(int, string) (string, error)
}

type UserHandler struct {
	UserService UserService
	Logger      logger.Logger
	Sessions    SessionManager
}

func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	body, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		uh.Logger.Errorw("can`t read body of request",
			"err:", err.Error())
		http.Error(w, "bad reg data", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, user)
	if err != nil {
		uh.Logger.Infow("can`t unmarshal register form",
			"err:", err.Error())
		http.Error(w, "bad reg data", http.StatusBadRequest)
		return
	}

	_, err = govalidator.ValidateStruct(user)
	if err != nil {
		uh.Logger.Infow("can`t validate register form",
			"err:", err.Error())
		http.Error(w, "bad reg data", http.StatusBadRequest)
		return
	}

	if user.Role != "user" {
		uh.Logger.Infow("can`t register not user role",
			"err:", err.Error())
		http.Error(w, "bad reg data", http.StatusBadRequest)
		return
	}

	user.ID, err = uh.UserService.CreateUser(*user)
	if err != nil {
		uh.Logger.Infow("can`t create user (username is already used)",
			"err:", err.Error())
		http.Error(w, "username is already used", http.StatusBadRequest)
		return
	}

	token, err := uh.Sessions.CreateSession(user.ID, "user")
	if err != nil {
		uh.Logger.Errorw("can`t create session",
			"err:", err.Error())
		http.Error(w, "can`t make session", http.StatusInternalServerError)
		return
	}

	tokenForm := &TokenForm{token}
	resp, err := json.Marshal(tokenForm)

	if err != nil {
		uh.Logger.Errorw("can`t marshal session token",
			"err:", err.Error())
		http.Error(w, "can`t make session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_, err = w.Write(resp)
	if err != nil {
		uh.Logger.Errorw("can`t write response",
			"err:", err.Error())
		http.Error(w, "can`t write response", http.StatusInternalServerError)
		return
	}
}

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	regForm := &loginForm{}

	body, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		uh.Logger.Errorw("can`t read body of request",
			"err:", err.Error())
		http.Error(w, "bad reg data", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, regForm)
	if err != nil {
		uh.Logger.Infow("can`t unmarshal register form",
			"err:", err.Error())
		http.Error(w, "bad reg data", http.StatusBadRequest)
		return
	}

	_, err = govalidator.ValidateStruct(regForm)
	if err != nil {
		uh.Logger.Infow("can`t validate register form",
			"err:", err.Error())
		http.Error(w, "bad reg data", http.StatusBadRequest)
		return
	}

	user, err := uh.UserService.GetUserByLoginAndPassword(regForm.Login, regForm.Password)
	if err != nil {
		uh.Logger.Infow("can`t get user by login and password",
			"err:", err.Error())
		http.Error(w, "can`t login", http.StatusUnprocessableEntity)
		return
	}

	token, err := uh.Sessions.CreateSession(user.ID, user.Role)
	if err != nil {
		uh.Logger.Errorw("can`t create session",
			"err:", err.Error())
		http.Error(w, "can`t make session", http.StatusInternalServerError)
		return
	}

	tokenForm := &TokenForm{token}
	resp, err := json.Marshal(tokenForm)

	if err != nil {
		uh.Logger.Errorw("can`t marshal session token",
			"err:", err.Error())
		http.Error(w, "can`t make session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(resp)
	if err != nil {
		uh.Logger.Errorw("can`t write response",
			"err:", err.Error())
		http.Error(w, "can`t write response", http.StatusInternalServerError)
		return
	}
}
