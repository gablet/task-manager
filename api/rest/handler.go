package rest

import (
	"gabyRepos/task-manager/model"
	"gabyRepos/task-manager/service/tasks"
	"git.detectify.net/go/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	Router  *mux.Router
	Manager tasks.Manager
}

func NewHandler(router *mux.Router) *Handler {
	return &Handler{Router: router}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task *model.Task

	decoded := json.NewDecoder(r.Body)
	err := decoded.Decode(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = h.Manager.CreateTask(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if len(id) == 0 {
		w.WriteHeader(http.StatusBadRequest)
	}

	err := h.Manager.RemoveTask(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) CreateRoutes() []Route {
	return []Route{
		{
			Name:    "CreateTask",
			Method:  http.MethodPost,
			Pattern: "/task",
			Handler: http.HandlerFunc(h.CreateTask),
		},
		{
			Name:    "DeleteTask",
			Method:  http.MethodDelete,
			Pattern: "/task/{id}",
			Handler: http.HandlerFunc(h.DeleteTask),
		},
	}
}
