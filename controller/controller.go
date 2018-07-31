package controller

import (
	"encoding/json"
	"fmt"
	"apiToDo/model"
	"net/http"
)

//DB ...
var DB = model.DB

// Post ...
type Post struct {
	Titulo string `json:"titulo"`
	Texto  string `json:"texto"`
}

//Message ...
type Message  struct {
	Mensagem string `json:"mensagem"`
}
// Posts ...
type Posts []Post

//Messages ...
type Messages []Message
//Insert ...
func Insert(w http.ResponseWriter, r *http.Request) {
	titulo := r.FormValue("titulo")
	texto := r.FormValue("texto")

	_, err := DB.Query("INSERT INTO tb_todo(titulo, texto) VALUES('" + titulo + "','" + texto + "')")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Post inserido com sucesso")
	messages:=Messages{
		Message{Mensagem: "Post inserido com sucesso"},
	}
	json.NewEncoder(w).Encode(messages)
}

// Delete ...
func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	_, err := DB.Query("DELETE FROM tb_todo WHERE id = " + id )
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Post deletado com sucesso")
	messages:=Messages{
		Message{Mensagem: "Post deletado com sucesso"},
	}
	json.NewEncoder(w).Encode(messages)
}

// //Select ...
// func Select(w http.ResponseWriter, r *http.Request) {
// 	model.Select(w, r)
// }
