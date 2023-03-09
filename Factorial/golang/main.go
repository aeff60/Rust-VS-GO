package main

import (
	"fmt"
)

func main() {
	fmt.Println("โปรแกรมคำนวน Factorial")

	var n uint64

	fmt.Println("กรุณากรอกจำนวนเต็มบวก: ")
	fmt.Scan(&n)

	var result uint64 = 1

	for i := uint64(1); i <= n; i++ {
		result *= i
	}

	fmt.Printf("%d! = %d", n, result)
}
