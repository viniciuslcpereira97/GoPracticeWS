package routes

import (
    "fmt"
    "html"
    "net/http"
    "github.com/gorilla/mux"
)

var Router *mux.Router

func InitRouter() {
    if Router == nil {
        Router = mux.NewRouter()
    }
}

func AddRoutes(route string, methods string) {
    InitRouter()
    Router.HandleFunc(route, handler).Methods(methods)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Serve() {
    AddRoutes("/teste","GET")
    http.ListenAndServe(":8888", Router)
}
