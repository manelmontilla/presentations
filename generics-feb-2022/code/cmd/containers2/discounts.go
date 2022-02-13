package main

// START OMIT
func hasDisccounts(c Cache, name string) bool {
	for _, n := range c.Keys() {
		nstring, ok := n.(string) // HL
		if !ok {
			panic("what can I do?") // HL
		}
		if nstring != name {
			continue
		}
		items := c.Get(n)
		intItems, ok := items.(int)
		if !ok {
			panic("what can I do?") // HL
		}
		if intItems > 1 {
			return true
		}
	}
	return false
} // END OMIT
