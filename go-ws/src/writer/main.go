package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "html/template"
    "db"
)

type ServerCTX struct {
    idgen     IdGenerator
    cache     Cache     
    dbInst    db.DB
}

type writeOk struct {
    Url string
    Id string
}

func writeHandler(ctx ServerCTX) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
    	r.ParseForm()
	    u := r.Form.Get("url")
        fmt.Println("Url is " + u)
        v := ctx.cache.getValue(u)
        fmt.Println("value " + v)
	    if v == "" {
            fmt.Println("Trying to get id " + v)
            v = ctx.idgen.getId()
            fmt.Println("Got value " + v)
            ctx.dbInst.AddToDB(u, v)
    	    ctx.cache.addKV(u, v)
        }
        t, _ := template.ParseFiles("writeok.html")
        resp := writeOk{Url: u, Id: v}
        t.Execute(w, resp)
    }
    return http.HandlerFunc(fn)
}

func mainHandler(ctx ServerCTX) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        data, _ := ioutil.ReadFile("./main.html")
        w.Write(data)
    }
    return http.HandlerFunc(fn)
}

func main() {

    g := GetIdGenerator()
    c := CreateCache()
    d := db.CreateDB()
    serverCTX :=  ServerCTX{idgen : g, cache: c, dbInst:d}
    
    mux := http.NewServeMux()
    mux.Handle("/", mainHandler(serverCTX))
    mux.Handle("/main", mainHandler(serverCTX))
    mux.Handle("/write", writeHandler(serverCTX))
    http.ListenAndServe(":8000", mux)

}
