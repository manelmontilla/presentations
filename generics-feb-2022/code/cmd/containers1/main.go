package main

// START OMIT
type Cache struct {
	cache map[interface{}]interface{}
}

func (c Cache) Add(key, value interface{}) {
	c.cache[key] = value
}

func (c Cache) Get(key interface{}) interface{} {
	return c.cache[key]
}

func main() {
	c := Cache{cache: make(map[interface{}]interface{})}
	c.Add("a", "b")      // HL
	c.Add([]int{1}, "a") // HL
}

//END OMIT
