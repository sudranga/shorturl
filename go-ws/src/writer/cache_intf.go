package main

type Cache interface {
    init() 
    addKV(url string, sf string)
    getValue(v string) string
}

func CreateCache() Cache{
    v := &redisCache{}
    v.init()
    return v
}
