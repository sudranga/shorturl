package main

import "os"

type Cache interface {
    init(connString string) 
    addKV(url string, sf string)
    getValue(v string) string
}

func CreateCache() Cache{
    redisSvcHost := os.Getenv("REDIS_SERVICE_HOST")
    redisSvcPort := os.Getenv("REDIS_SERVICE_PORT")
    if redisSvcHost == "" {
        redisSvcHost = "redis"
    }

    if redisSvcPort == "" {
        redisSvcPort = "6379"
    }
    redisSvc := redisSvcHost + ":" + string(redisSvcPort)

    v := &redisCache{}
    v.init(redisSvc)
    return v
}
