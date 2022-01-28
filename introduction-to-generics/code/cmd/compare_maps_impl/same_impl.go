package main

func main() {

}

// START OMIT
// Same implementation, only different types in the headers. // HL
func StringSetsEqual(m1, m2 map[string]struct{}) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func StringsMapsEqual(m1, m2 map[string]string) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}

// END OMIT
