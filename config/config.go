package config

import (
    _"fmt"
    "os"
    "strings"
    "reflect"
    "encoding/json"
)

const FILE_PATH = "./config/config.json"

type Config struct {
    Db string
}

var configuration Config

func SetUp() {

    config_file, _err := os.Open(FILE_PATH)

    if _err != nil {
        panic(_err)
    }

    decoder := json.NewDecoder(config_file)
    _err = decoder.Decode(&configuration)

    if _err != nil {
        panic(_err)
    }

}

func GetConfig(conf string) (conf_field interface {}){
    
    SetUp()

    valr := reflect.ValueOf(configuration)
    conf_field = reflect.Indirect(valr).FieldByName(strings.Title(conf)).Interface()

    return conf_field

}