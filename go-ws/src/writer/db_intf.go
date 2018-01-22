package main

type DB interface {
    initDB()
    addToDB(key string, val string) 
}

func CreateDB() DB {
    z := &mySQLDB{}
    z.initDB()
    return z
}
