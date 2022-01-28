package genericsp

// START OMIT
func StringsMapsEqual(m1, m2 map[string]string) bool { // HL
	if len(m1) != len(m2) {
		return false
	}
	for k1, v1 := range m1 {
		v2, ok := m2[k1]
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func IntStringsMapsEqual(m1, m2 map[int]string) bool { // HL
	if len(m1) != len(m2) {
		return false
	}
	for k1, v1 := range m1 {
		v2, ok := m2[k1]
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func IntsMapsEqual(m1, m2 map[int]int) bool { // HL
	if len(m1) != len(m2) { // HLEQ
		return false // HLEQ
	}
	for k1, v1 := range m1 { // HLEQ
		v2, ok := m2[k1]     // HLEQ
		if !ok || v1 != v2 { // HLEQ
			return false // HLEQ
		} // HLEQ
	} // HLEQ
	return true // HLEQ
}

func StringSetsEqual(m1, m2 map[string]struct{}) bool { // HL
	if len(m1) != len(m2) { // HLEQ
		return false // HLEQ
	} // HLEQ
	for k1, v1 := range m1 { // HLEQ
		v2, ok := m2[k1]     // HLEQ
		if !ok || v1 != v2 { // HLEQ
			return false // HLEQ
		} // HLEQ
	} // HLEQ
	return true // HLEQ
} // HLEQ

// END OMIT
