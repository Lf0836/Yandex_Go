package main

func CountingSort(contacts []string) map[string]int {
	counts := make(map[string]int)
	for _, contact := range contacts {
		counts[contact]++
	}
	return counts
}
