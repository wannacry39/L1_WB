package main

import "fmt"

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

// 1. Глобальная переменная justString. Глобальные переменные приводят к неконтролируемой области видимости и усложнению читаемости кода.
// 2. Не возвращаем значение из функции someFunc(), а просто присваиваем полученный срез переменной justString. Может привести к тому, что
// указатель не будет удален GC и будет занимать место в памяти.

func main() {
	justString := someFunc()
	fmt.Printf("%T", justString)
	fmt.Println(justString)
}

func someFunc() string {
	v := createHugeString(1 << 10)
	return v[:100]
}

func createHugeString(n int) string {
	return string(make([]byte, n))
}
