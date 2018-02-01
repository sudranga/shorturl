package db

import "os"

type DB interface {
    InitDB(conn map[string]string)
    AddToDB(key string, val string) 
    ReadFromDB(id string) (string, error)
}

func CreateDB() DB {
    params := make(map[string]string)
    tmp := os.Getenv("MYSQL_SERVICE_HOST")
    if tmp == "" {
        tmp = "mysql"
    }
    params["Username"] = "root"
    params["Password"] = "root"
    params["Host"] = tmp
    params["Db"] = "test"
    
    z := &mySQLDB{}
    z.InitDB(params)
    return z
}
