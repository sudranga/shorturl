package db

import (
	"database/sql"
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"
)

type mySQLDB struct {
    db    *sql.DB
}

func (d *mySQLDB) InitDB(params map[string]string) {
    
    db, err := sql.Open("mysql", params["Username"]+":" + params["Password"] + "@tcp(" + params["Host"] + ")/"+params["Db"])
    if err != nil {
        panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
    }

    // Open doesn't open a connection. Validate DSN data:
    err = db.Ping()
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    d.db = db
}

func (d *mySQLDB) ReadFromDB(id string) string {
    stmt, err := d.db.Prepare("SELECT * FROM testtab WHERE id= ?")
    if err != nil {
        panic(err.Error())
    }
    var tmp1 int
    var tmp2 int
    var url string

    err = stmt.QueryRow(id).Scan(&tmp1, &tmp2, &url)
    if err != nil {
        panic(err.Error())
    }
    
    return url
}

func (d *mySQLDB) AddToDB(key string, val string) {
    stmt, err := d.db.Prepare("INSERT INTO testtab (url, id) VALUES(?,?)") // ? = placeholder
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    _, err = stmt.Exec(key, val)
    if err != nil {
        panic(err.Error())
    }
}
