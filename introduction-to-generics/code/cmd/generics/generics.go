package generics

func MapsEqual[T comparable, K comparable](m1, m2 map[T]K) bool {
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

func compareStringMaps() {
	m1 := map[string]string{}
	m2 := map[string]string{}
	equal := MapsEqual[string, string](m1, m2)
	println(equal)
}

func compareIntStringMaps() {
	m1 := map[int]string{}
	m2 := map[int]string{}
	equal := MapsEqual[int, string](m1, m2)
	println(equal)
}

func compareStringMapsInference() {
	firstM := map[string]string{}
	secondM := map[string]string{}
	equal := MapsEqual(firstM, secondM)
	println(equal)
}
