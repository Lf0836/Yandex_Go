package main

func SwapKeysAndValues(m map[string]string) map[string]string {
	swapped := make(map[string]string)
	for key, value := range m {
		swapped[value] = key
	}
	return swapped
}
