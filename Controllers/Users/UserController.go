package usercontroller

import (
    "fmt"
    "net/http"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "All Users")
}