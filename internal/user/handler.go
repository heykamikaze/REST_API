package user

import (
	"awesomeProject/internal/handlers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var _ handlers.Handler = &handler{} //если че то поменяется ошибка сразу будет видна здесь

const (
	usersURL = "/users"
	userURL  = "/user/:uuid"
)

type handler struct {
} //структура хендлера реализует интерфейс хендлер logger service

func NewHandler() handlers.Handler {
	return &handler{} //создаем струткуру но возвращаем интерфейс, структура неэкспортируемая
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.POST(usersURL, h.CreateUser)
	router.GET(userURL, h.GetUserByUUID)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("this is list of users"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200) //if everythjing is fine
	w.Write([]byte("this is list of users by uuid"))
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(201)
	w.Write([]byte("this is create user"))
}
func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is update of users"))
}
func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is partially update of users"))
}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is delete user"))
}
