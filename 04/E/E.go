package main

import "fmt"

func main() {
	var n, ans int
	fmt.Scan(&n)
	ans = 0
	if n < 0 {
		fmt.Print("Некорректный ввод")
	} else {
		for i := 1; i < n+1; i += 2 {
			ans += i
		}
		fmt.Print(ans)
	}
}
