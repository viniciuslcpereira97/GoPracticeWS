package routes

import (
    "net/http"
    "github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
    return
}

func Router() (r *mux.Router){
    r = mux.NewRouter()
    r.HandleFunc("/teste", handler).Methods("GET")

    return r
}