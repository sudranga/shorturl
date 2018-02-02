package cache

import "os"

type Cache interface {
    Init(connString string) 
    AddKV(url string, sf string)
    GetValue(v string) string
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
    v.Init(redisSvc)
    return v
}
