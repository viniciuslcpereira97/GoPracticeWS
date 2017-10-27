package main

import (
    "fmt"
    "practice-ws/DB/Models"
)

func main() {
    
    condition := 
        map[string] interface {} {
            "Age": 20, 
        }

    fields := 
        map[string] interface {} {
            "_id": 0,
        }

    user := models.Find(condition, fields)

    fmt.Println(user)

}