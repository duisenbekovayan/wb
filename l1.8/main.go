package main

import "fmt"

// Установить i-й бит в 1
func setBit(x int64, i uint) int64 {
	return x | (1 << i)
}

// Установить i-й бит в 0
func clearBit(x int64, i uint) int64 {
	return x &^ (1 << i)
}

func main() {
	var n int64 = 5 // 0101
	var i uint = 0  // номер бита (считаем с 0 справа)

	fmt.Printf("Исходное: %b (%d)\n", n, n)

	n1 := setBit(n, i)
	fmt.Printf("Установить бит %d в 1: %b (%d)\n", i, n1, n1)

	n2 := clearBit(n, i)
	fmt.Printf("Установить бит %d в 0: %b (%d)\n", i, n2, n2)
}
