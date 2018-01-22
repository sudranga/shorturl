package main

import (
    "github.com/samuel/go-zookeeper/zk"
    _ "github.com/curator-go/curator"
    "time"
    "fmt"
    "strconv"
)

type zkInst struct {
    conn    *zk.Conn
    currId  int64
    maxId   int64
}

func must(err error) {
    if err != nil {
        panic(err)
    }
}

func (z *zkInst) init() {
    z.currId = 0
    z.maxId = 0
    var servers []string
    servers = append(servers, "zookeeper")
    conn, _, err := zk.Connect(servers, time.Second)
    if err != nil {
        panic(err)
    }
    z.conn = conn
    // defer z.conn.Close()

    flags := int32(0)
    acl := zk.WorldACL(zk.PermAll)

    path, err := conn.Create("/ctr", []byte("0"), flags, acl)
    if err != nil && err != zk.ErrNodeExists {
        panic(err)
    }
    fmt.Printf("create: %+v\n", path)

    /*retryPolicy := curator.NewExponentialBackoffRetry(time.Second, 3, 15*time.Second)
    client := curator.NewClient("connString", retryPolicy)
    client.Start()
    defer client.Close()*/
}

func (z *zkInst) getId() string {
    if z.currId == z.maxId {
        for {
            fmt.Println("get Loop")
            data, stat, err := z.conn.Get("/ctr")
            must(err)
            fmt.Printf("get:    %+v %+v\n", string(data), stat)
            z.currId, _ = strconv.ParseInt(string(data), 10, 64)
            z.maxId = z.currId + 2048
            buf := []byte(strconv.FormatInt(z.maxId, 10))
            stat, err = z.conn.Set("/ctr", buf , stat.Version)
            if err == nil {
                break
            }
            if err == zk.ErrBadVersion {
                continue
            }
            must(err)
            fmt.Printf("set:    %+v\n", stat)
        }
    }
    tmpId := z.currId
    z.currId = z.currId + 1
    return strconv.FormatInt(int64(tmpId), 16)
}

