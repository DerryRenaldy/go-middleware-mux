package main

import (
	"coba-stripe-api/middleware"
	"coba-stripe-api/student"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func WelcomeMessage(w http.ResponseWriter, r *http.Request) {
	OutputJSON(w, map[string]string{"message": "welcome to basic auth practice, this no need the authorization"})
}

func AuthMessage(w http.ResponseWriter, r *http.Request) {
	OutputJSON(w, map[string]string{"message": "to access this endpoint you need authorization"})
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {

	//switch r.Method {
	//case http.MethodGet:
	//	w.Write([]byte("Method Get"))
	//
	//case http.MethodPost:
	//	w.Write([]byte("Method Post"))
	//default:
	//	http.Error(w, "", http.StatusBadRequest)
	//}

	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, student.SelectStudent(id))
		return
	}
	OutputJSON(w, student.GetStudent())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/welcome", WelcomeMessage).Methods(http.MethodGet)

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/student", ActionStudent).Methods(http.MethodGet)
	api.HandleFunc("/auth", AuthMessage).Methods(http.MethodGet)
	api.Use(middleware.MiddlewareAllowOnlyGet, middleware.MiddlewareAuth)
	//server := new(http.Server)
	server := http.Server{
		Addr:    ":9000",
		Handler: r,
	}

	fmt.Println("server started at localhost:9000")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
