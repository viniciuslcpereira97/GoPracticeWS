package usercontroller

import (
    "reflect"
    "fmt"
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "GoPracticeWS/DB/Models/Users"
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
    data := r.Form
    user := make(map[string] interface {})
    fmt.Print(reflect.TypeOf(data["Age"][0]))
    for key, _ := range data {
        value := data[key][0]
        if intVal, err := strconv.Atoi(value); err == nil {
            user[key] = intVal
        } else {
            user[key] = value
        }
    }
    users.Create(user)
}