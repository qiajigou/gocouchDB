package gocouchDB

/*
A simple Memory cache reduce http requests
You can read ICache interface and use redis as well :)
*/

type MemCache struct {
    Hash map[string]interface{}
}

func NewMemCache() *MemCache{
    cl := new(MemCache)
    cl.Hash = map[string]interface{} {}
    return cl
}

func (cl *MemCache)Get(key string, defaultValue interface{})(value interface{}){
    value, ok := cl.Hash[key]

    if !ok {
        return defaultValue
    } else {
        return value
    }
}

func (cl *MemCache)Set(key string, value interface{})(ret bool){
    cl.Hash[key] = value
    return true
}

func (cl *MemCache)Delete(key string)(ret bool){
    delete(cl.Hash, key)
    return true
}