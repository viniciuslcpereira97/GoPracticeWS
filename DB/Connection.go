package db

import (
    "gopkg.in/mgo.v2"
    "practice-ws/config"
)

// Connect to Mongo Database
func Connect() (* mgo.Database) {

    db_host := config.GetConfig("db_host").(string)
    db_name := config.GetConfig("db_name").(string)

    session, err := mgo.Dial(db_host)
    
    if err != nil {
        panic(err)
    }

    return session.DB(db_name)

}