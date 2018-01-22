package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "html/template"
)

type ServerCTX struct {
    idgen IdGenerator
    cache Cache     
    db    DB
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
            ctx.db.addToDB(u, v)
    	    ctx.cache.addKV(u, v)
        }
        t, _ := template.ParseFiles("writeok.html")
        resp := writeOk{Url: u, Id: v}
        t.Execute(w, resp)
    }
    return http.HandlerFunc(fn)
}

func helloHandler(ctx ServerCTX) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        data, _ := ioutil.ReadFile("./main.html")
        w.Write(data)
    }
    return http.HandlerFunc(fn)
}

func main() {

    g := GetIdGenerator()
    c := CreateCache()
    d := CreateDB()
    serverCTX :=  ServerCTX{idgen : g, cache: c, db:d}
    
    mux := http.NewServeMux()
    mux.Handle("/", helloHandler(serverCTX))
    mux.Handle("/write", writeHandler(serverCTX))
    http.ListenAndServe(":8000", mux)

}
