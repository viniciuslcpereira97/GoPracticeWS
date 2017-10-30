package db

import (
    "reflect"
    "gopkg.in/mgo.v2"
    "practice-ws/config"
)

type DatabaseStruct struct {
    session *mgo.Database
}

var DB_CONN = DatabaseStruct {}

// Get Db instance
func GetInstance() (*mgo.Database) {
    if reflect.DeepEqual(DatabaseStruct {}, DB_CONN) {
        Connect()
    }

    return DB_CONN.session
}

// Connect to Mongo Database
func Connect() {
    db_host := config.GetConfig("db_host").(string)
    db_name := config.GetConfig("db_name").(string)

    session, err := mgo.Dial(db_host)
    
    if err != nil {
        panic(err)
    }

    DB_CONN.session = session.DB(db_name)
}