package main

import (
	"fmt"
	"wbtechl1tasks/taskutils"
)

const MIN_SEG_L = 64 // константа взята с потолка и по-хорошему должна вариироваться в зав. от железа

/*

Чтож, у нас имеется алгоритм за O(n) который надо как-то заумно горутинизировать
Так, чтобы это было асимптотически хорошо и в то же время мы не испытывали сильную просадку
Из-за старта-финиша горутин

Для этого используем комбинированный алгоритм, учитывающий разницу констант
Можно провести аналогию с IntroSort - сортировка, которая использует
сортировку вставкам на маленьких подотрезках в рамках быстрой сортировки

Мы будем рекурсивно углубляться в отрезок, пока не достигнем отрезков длиной
НУ ДОПУСТИМ 64 и на них мы будем считать сумму линейно
А дальше будем действовать как будто это дерево отрезков,
суммируя левый-правый отрезки

*/

type Segmentreeno struct {
	arr []int
	n   int
}

func newTreeno(arr []int) Segmentreeno {
	/*
		Выделяем память под структуру, копируем в неё массив (поверхностно)
		и возвращаем эту структуру
	*/
	s := new(Segmentreeno)
	s.n = len(arr)
	s.arr = make([]int, len(arr))
	copy(s.arr, arr)
	return *s
}

func getMid(left int, right int) int {
	return (left + right) / 2 // ср. арифметическое
}

func (s *Segmentreeno) sumSquareSegment(left int, right int, c chan<- int) {
	/*
		Если длина отрезка меньше чем MIN_SEG_L считаем сумму линейно
		 чтобы не тратить время на жиненный цикл горутин
		 Кидаем сумму в канал и завершаем функцию
	*/
	if right-left < MIN_SEG_L {
		sum := 0
		for i := left; i < right; i++ {
			sum += s.arr[i] * s.arr[i]
		}
		c <- sum
		return
	}

	/*
		В ином случае, углубляемся в рекурсию: делим подотрезок на равные отрезки
		и запускаемся от них. Ждём результаты из канала и возвращаем их сумму
	*/

	sums := make(chan int, 2)
	m := getMid(left, right)
	go s.sumSquareSegment(left, m, sums)
	go s.sumSquareSegment(m, right, sums)
	c <- (<-sums + <-sums)
}

func (s *Segmentreeno) SumSquares() int {
	/*
	 Входная точка, возвращающая итоговую сумму. Запускает рекурсивную функцию
	 На отрезке от начала и до конца
	*/
	c := make(chan int)
	go s.sumSquareSegment(0, s.n, c)
	return <-c
}

func main() {
	var arr []int

	taskutils.ReadSlice(&arr)
	treeno := newTreeno(arr)
	s := treeno.SumSquares()
	fmt.Println(s)
}
