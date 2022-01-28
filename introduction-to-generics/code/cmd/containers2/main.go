package main

import "fmt"

type Cache struct {
	cache map[interface{}]interface{}
}

func (c Cache) Add(key, value interface{}) {
	c.cache[key] = value
}

func (c Cache) Get(key interface{}) interface{} {
	v := c.cache[key]
	return v
}

func (c Cache) Keys() []interface{} {
	var keys []interface{}
	for k := range c.cache {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	// START OMIT
	c := Cache{cache: make(map[interface{}]interface{})}
	c.Add("manel", 2) // HL
	if hasDisccounts(c, "manel") {
		fmt.Println("applying disccount to manel")
	}
	// END OMIT
}
