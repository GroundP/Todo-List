package main

import (
	"net/http"
	"sort"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rd *render.Render
var todoMap map[int]Todo
var lastID int = 0

type Todo struct {
	ID        int    `json:"id,mitempty"`
	Name      string `json:"name"`
	Completed bool   `json:"completed,mitempty"`
}

func MakeWebHandler() http.Handler {
	todoMap = make(map[int]Todo)
	mux := mux.NewRouter()
	mux.Handle("/", http.FileServer(http.Dir("public")))
	mux.HandleFunc("/todos", GetTodoListHandler).Methods("GET")
	mux.HandleFunc("/todos", PostTodoHandler).Methods("POST")
	mux.HandleFunc("/todos/{id:[0-9]+}", RemoveTodoHandler).Methods("DELETE")
	mux.HandleFunc("/todos/{id:[0-9]+}", UpdateTodoHandler).Methods("PUT")
	return mux
}

type Todos []Todo
func (t Todos) Len() int {
	return len(t)
}

func (t Todos) Swap(i,j int){
	t[i], t[j] = t[j], t[i]
}

func (t Todos) Less(i,j int) bool {
	return t[i].ID > t[j].ID
}

func GetTodoListHandler(w http.ResponseWriter r *http.Request){
	list := make(Todos, 0)
	for _, todo := range todoMap {
		list = append(list, todo)
	}
	sort.Sort(list)
	rd.JSON(w, http.StatusOK, list)
}

