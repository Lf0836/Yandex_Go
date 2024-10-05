package main

func DeleteLongKeys(m map[string]int) map[string]int {
	newMap := make(map[string]int)
	for key, value := range m {
		if len(key) >= 6 {
			newMap[key] = value
		}
	}
	return newMap
}
