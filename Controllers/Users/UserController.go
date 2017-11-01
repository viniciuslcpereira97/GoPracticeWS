package usercontroller

import (
    "fmt"
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "practice-ws/DB/Models/Users"
)

// Get All users
func All(w http.ResponseWriter, r *http.Request) {
    users, _ := json.Marshal(users.All())
    fmt.Fprint(w, string(users))
}

// Get Users By Age
func ByAge(w http.ResponseWriter, r *http.Request) {
    user_age, _ := strconv.ParseInt(mux.Vars(r)["age"], 10, 64)
    condition := map[string] int64 {
        "age": user_age,
    }
    users, _ := json.Marshal(users.Where(condition))
    fmt.Fprint(w, string(users))
}

// Create new user
func CreateNew(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    data := r.PostForm
    user := make(map[string] interface {})
    for key, _ := range data {
        user[key] = r.PostFormValue(key) 
    }
    users.Create(user)
}