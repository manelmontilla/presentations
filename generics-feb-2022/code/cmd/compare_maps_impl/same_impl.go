package main

// START OMIT
func StringSetsEqual(m1, m2 map[string]struct{}) bool {
	if len(m1) != len(m2) {
		return false
	} // HL
	for k1, v1 := range m1 { // HL
		v2, ok := m2[k1]     // HL
		if !ok || v1 != v2 { // HL
			return false // HL
		} // HL
	} // HL
	return true // HL
} // HL

func StringsMapsEqual(m1, m2 map[string]string) bool {
	if len(m1) != len(m2) {
		return false
	} // HL
	for k1, v1 := range m1 { // HL
		v2, ok := m2[k1]     // HL
		if !ok || v1 != v2 { // HL
			return false // HL
		} // HL
	} // HL
	return true // HL
}

// END OMIT
