package models

import (
    "practice-ws/DB"
)

var collection = db.Connect().C("users")

type User struct {
    Name string
    Age int
    Status string
}

func Create(user interface {}) {
    collection.Insert(&user)
}

func Find(conditions interface {}, fields interface {}) (user map[string] interface {}) {
    collection.Find(conditions).Select(fields).One(&user)
    return user
}