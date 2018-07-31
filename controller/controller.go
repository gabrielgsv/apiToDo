package controller

import (
	"apiToDo/model"

	// "encoding/json"
	"net/http"
)

model.TryConn()

//Post ...
type Post struct {
	Titulo string `json:"titulo"`
	Texto  string `json:"texto"`
}

//Posts ...
type Posts []Post

//Insert ...
func Insert(w http.ResponseWriter, r *http.Request) {
	titulo := r.FormValue("titulo")
	texto := r.FormValue("texto")

	model.Create(titulo, texto)
}

//Delete ...
func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	model.Delete(id)
}

//Select ...
func Select(w http.ResponseWriter, r *http.Request) {
	model.Select(w, r)
}
