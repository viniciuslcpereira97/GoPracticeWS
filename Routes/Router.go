package routes

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "practice-ws/Controllers/Users"
)

// Route struct
type Route struct {
    Name string
    Path string
    Methods string
    Handler http.HandlerFunc
}

// Array of Routes
var routes []Route

// Router instance
var Router *mux.Router

// Initialize new Router instance
func InitRouter() {
    if Router == nil {
        Router = mux.NewRouter()
    }
}

// Register POST route
func Post(name string, route string, handler http.HandlerFunc) {
    new_route := Route {
        name,
        route,
        "POST",
        handler,
    }

    routes = append(routes, new_route)
}

// Register new GET route
func Get(name string, route string, handler http.HandlerFunc) {
    new_route := Route {
        name, 
        route,
        "GET",
        handler,
    }

    routes = append(routes, new_route)
}

// Handle all registered routes before provide Router
func LoadRoutes() {
    InitRouter()
    AddRoutes()
    for _, route := range routes {
        fmt.Printf("route_name: %s  -->  route_pattern: %s\n", route.Name, route.Path)
        Router.HandleFunc(route.Path, route.Handler).Methods(route.Methods)
    }
}

// Loads all Routes
// Start to Listen and Serve
func Serve() {
    LoadRoutes()
    http.ListenAndServe(":8888", Router)
}

// Temporary function to add routes
// TODO: Separate in files of specific services
func AddRoutes() {

    //Users
    Get("users.index", "/users", usercontroller.All)
    Get("users.byAge", "/users/age/{age}", usercontroller.ByAge)
}