package users

import (
    "practice-ws/DB"
    "github.com/mitchellh/mapstructure"
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
func Find(conditions interface {}, fields ...interface {}) (user User) {
    var result interface {}
    collection.Find(conditions).Select(fields).One(&result)
    mapstructure.Decode(result, &user)

    return user
}

// List of Users that satisfies conditions
func Where(conditions interface {}, fields ...interface {}) (users []User) {
    var result []interface{}
    collection.Find(conditions).Select(fields).All(&result)
    mapstructure.Decode(result, &users)

    return users
}

func ById(user_id string) (user User) {
    var result interface {}
    condition := map[string] string {
        "_id": user_id,
    }
    collection.Find(condition).Select(nil).One(&result)
    mapstructure.Decode(result, &user)

    return user
}

func All() (users []User) {
    var result []interface{}
    collection.Find(nil).All(&result)
    mapstructure.Decode(result, &users)

    return users
}