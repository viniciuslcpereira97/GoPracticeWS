package main

import (
    "fmt"
    "net/http"
    "practice-ws/Routes"
)

func main() {
    routes.Get("say.hello", "/hello", say_hello)
    routes.Serve()
}

func say_hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello World!")
}