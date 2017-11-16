package routes

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/viniciuslcpereira97/GoPracticeWS/Controllers/Users"
)

// Route struct
type Route struct {
    Name string
    Path string
    Methods string
    Handler http.HandlerFunc
}

// Array of Routes
var routes = make(map[string] Route)

// Router instance
var Router *mux.Router

// Initialize new Router instance
func InitRouter() {

    if Router == nil {
        Router = mux.NewRouter()
    }

}

func checkIfRouteIsDuplicated(name string) {

    if _, route_exists := routes[name]; route_exists {
        panic("route " + name + " exists at routes list")
    }

}

// Register POST route
func Post(name string, route string, handler http.HandlerFunc) {

    checkIfRouteIsDuplicated(name)

    new_route := Route {
        name,
        route,
        "POST",
        handler,
    }

    routes[name] = new_route
}

// Register new GET route
func Get(name string, route string, handler http.HandlerFunc) {
    
    checkIfRouteIsDuplicated(name)

    new_route := Route {
        name, 
        route,
        "GET",
        handler,
    }

    routes[name] = new_route
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
    Post("users.createNew", "/users/new", usercontroller.CreateNew)

}