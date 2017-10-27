package models

import (
    "practice-ws/DB"
)

var collection = db.GetInstance().C("users")

// User struct
type User struct {
    Name string
    Age int
    Status string
}

// Insert new user at database
func Create(user interface {}) {
    collection.Insert(&user)
}

// Find User by conditions
func Find(conditions interface {}, fields interface {}) (user map[string] interface {}) {
    collection.Find(conditions).Select(fields).One(&user)
    return user
}