package cache

import (
    "github.com/go-redis/redis"
    "fmt"
)

type redisCache struct {
    client *redis.Client
}

func (c *redisCache) AddKV(url string, shortForm string) {
    fmt.Println("R: Add url: "+url+" and shortform: "+shortForm)
    c.client.Set(url, shortForm, 0)
}

func (c *redisCache) GetValue(url string) string {
    fmt.Println("R: Find Url: "+url)
    val, err := c.client.Get(url).Result()
    fmt.Println("R: Found value:  "+val)
    fmt.Println(err)
    if err == nil {
        return val
    } else {
        return ""
    }
}

func (c *redisCache) Init(redisSvc string) {
    client := redis.NewClient(&redis.Options{
        Addr:     redisSvc,
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    pong, err := client.Ping().Result()
    fmt.Println(pong, err)
    c.client = client
    // Output: PONG <nil>
}
