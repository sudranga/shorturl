package main

type IdGenerator interface {
    init()
    getId() string
}

func GetIdGenerator() IdGenerator {
    z := &zkInst{}
    z.init()
    return z
}
