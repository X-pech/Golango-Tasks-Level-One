package main

import "fmt"

func SetIBit(a *int64, pos int, val int64) {
	*a &= ^(1 << int64(pos))  // Выставляем этот бит в 0, очищаем его
	*a |= (val << int64(pos)) // Выставляем в него нужное значение
}

func main() {
	var a int64 = 98765678908768934
	fmt.Printf("%b\n", a)
	SetIBit(&a, 4, 1)
	SetIBit(&a, 5, 0)
	fmt.Printf("%b\n", a)

}
