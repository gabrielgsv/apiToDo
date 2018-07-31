package routes
import (
	"apiToDo/model"
	"apiToDo/controller"
	
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

//HandleFunc ...
func HandleFunc(){
	r := mux.NewRouter()

	model.TryConn()

	r.HandleFunc("/api/insert", controller.Insert).Methods("POST")
	r.HandleFunc("/api/delete", controller.Delete).Methods("DELETE")
	r.HandleFunc("/api/select", controller.Select).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))
}
