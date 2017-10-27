package db

import (
    "gopkg.in/mgo.v2"
    "practice-ws/config"
)

func Connect() (* mgo.Database) {

    db_host := config.GetConfig("db").(string)
    session, err := mgo.Dial(db_host)
    
    if err != nil {
        panic(err)
    }

    return session.DB("practice_db")

}