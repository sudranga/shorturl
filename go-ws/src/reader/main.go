package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "db"
    "html/template"
    "net/url"
    "cache"
)

type ServerCTX struct {
    dbInst      db.DB
    cacheInst   cache.Cache
}

type readOk struct {
    Url string
    Id string
}

func readHandler(ctx ServerCTX) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        urlString := r.URL.Path
        parsedUrl, err := url.Parse(urlString)
        if err != nil {
            http.Error(w, "Unable to read request", 400)
            return
        }
	    id := parsedUrl.Path[1:len(parsedUrl.Path)]
	    fmt.Println("Id is " + id)
        
        u := ctx.cacheInst.GetValue(id)
        fmt.Println("Cached val is " + u)
        
        if u == "" {
            u, err = ctx.dbInst.ReadFromDB(id)
            if err == nil {
                ctx.cacheInst.AddKV(id, u)
            }
        }

        if err != nil {
            http.NotFound(w, r)
            return
        } else {
            http.Redirect(w, r, u, 303)
        }
    }
    return http.HandlerFunc(fn)
}

func debugReadHandler(ctx ServerCTX) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
    	r.ParseForm()
	    id := r.Form.Get("id")
	    fmt.Println("Id is " + id)
        u, err := ctx.dbInst.ReadFromDB(id)
        if err != nil {
            http.NotFound(w, r)
        } else {
            t, _ := template.ParseFiles("readok.html")
            resp := readOk{Url: u, Id: id}
            t.Execute(w, resp)
        }
    }
    return http.HandlerFunc(fn)
}

func debugMainHandler(ctx ServerCTX) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        data, _ := ioutil.ReadFile("./main.html")
        w.Write(data)
    }
    return http.HandlerFunc(fn)
}

func main() {

    d := db.CreateDB()
    c := cache.CreateCache()
    serverCTX :=  ServerCTX{dbInst: d, cacheInst: c}
    
    mux := http.NewServeMux()
    mux.Handle("/main", debugMainHandler(serverCTX))
    mux.Handle("/read", debugReadHandler(serverCTX))
    mux.Handle("/", readHandler(serverCTX))
    http.ListenAndServe(":8080", mux)

}
