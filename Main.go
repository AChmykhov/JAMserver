package main


import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/lib/pq"
)


func main() {
	r := mux.NewRouter()
	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full password=check123"
	db, err := sql.Open("postgres", connStr)

	r.HandleFunc("/notification/{user}", NewNotification).Methods("POST")
	r.HandleFunc("/notification/{user}", ReadNotification).Methods("GET")
//	r.HandleFunc("/notification/{user}", DeleteNotification).Methods("DELETE")



	if err != nil {
		log.Fatal(err)
	}
}

//func DeleteNotification(writer http.ResponseWriter, request *http.Request) {
//
//}

func ReadNotification(writer http.ResponseWriter, request *http.Request) {

}

func NewNotification(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()


}

func insertNotification(user string, targetUser string, msg string, key string, symIv string)  {


}

func DeleteNotification(notificationId int){

}
