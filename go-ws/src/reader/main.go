package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "html/template"
)

type ServerCTX struct {
    db    DB
}

type readOk struct {
    Url string
    Id string
}

func readHandler(ctx ServerCTX) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
    	r.ParseForm()
	id := r.Form.Get("id")
	fmt.Println("Id is " + id)
        u := ctx.db.readFromDB(id)
        t, _ := template.ParseFiles("readok.html")
        resp := readOk{Url: u, Id: id}
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

    d := CreateDB()
    serverCTX :=  ServerCTX{db:d}
    
    mux := http.NewServeMux()
    mux.Handle("/", helloHandler(serverCTX))
    mux.Handle("/read", readHandler(serverCTX))
    http.ListenAndServe(":8080", mux)

}
