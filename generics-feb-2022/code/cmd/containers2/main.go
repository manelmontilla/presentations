package main

import "fmt"

// START OMIT
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
	c := Cache{cache: make(map[interface{}]interface{})}
	name := "manel"
	items := 2
	c.Add(name, items) // HL
	has := hasDisccounts(c, name)
	if has {
		fmt.Printf("apply disccount to %+v", has)
	}
}

func hasDisccounts(c Cache, name string) bool {
	for _, n := range c.Keys() {
		nstring, ok := n.(string) // HL
		if !ok {
			panic("what can I do?") // HL
		}
		if nstring == name {
			items := c.Get(n)
			intItems, ok := items.(int)
			if !ok {
				panic("what can I do?") // HL
			}
			if intItems > 1 {
				return true
			}
		}
	}
	return false
}

//END OMIT
