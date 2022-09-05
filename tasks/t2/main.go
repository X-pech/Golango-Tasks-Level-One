package main

import (
	"wbtechl1tasks/taskutils"
)

const MIN_SEG_L = 64 //константа взята с потолка и по-хорошему должна вариироваться в зав. от железа

/*

Чтож, у нас имеется алгоритм за O(n) который надо как-то заумно горутинизировать
Так, чтобы это было асимптотически хорошо и в то же время мы не испытывали сильную просадку
Из-за старта-финиша горутин

Для этого используем комбинированный алгоритм, учитывающий разницу констант
Можно провести аналогию с IntroSort - сортировка, которая использует
сортировку вставкам на маленьких подотрезках в рамках быстрой сортировки

Мы будем конкурентно разбивать подотрезки массива на право и лево,
рекурсивно спускаясь в них до тех пор, пока не достигнем отрезка меньше чем MIN_SEG_L
Таким образом глубина рекурсии будет O(log n) и примерно с такой асимптотикой
времени выполнения будет работать алгоритм, где в константе учитывается setup+teardown
O(2n) горутин (сумма ряда n + n/2 + n/4 + ... + 2 + 1)

Запуск горутин происходит из других горутин чтобы законкурентить и этот процесс
линейно это бы заняло O(n/segments) операций
*/

type Segmentreeno struct {
	arr          []int
	n            int
	segmentCount int
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

	/*
	 Считем количество отрезков
	*/
	s.segmentCount = s.n / MIN_SEG_L
	if s.n%MIN_SEG_L != 0 {
		s.segmentCount++
	}

	return *s
}

func getMid(left int, right int) int {
	return (left + right) / 2 // ср. арифметическое
}

func (s *Segmentreeno) cntSquareSegment(left int, right int, c chan<- int) {
	/*
		Если длина отрезка меньше чем MIN_SEG_L
		То каждый элемент на нём домножаем на самого себя
		В канал возвращаем количество элементов и завершаем функцию
	*/
	if right-left < MIN_SEG_L {
		for i := left; i < right; i++ {
			s.arr[i] *= s.arr[i]
		}
		c <- (right - left)
		return
	}

	/*
		Иначе разбиваем отрезок пополаем и запускаем горутины
	*/

	m := getMid(left, right)
	go s.cntSquareSegment(left, m, c)
	go s.cntSquareSegment(m, right, c)
}

func (s *Segmentreeno) CntSquares() {
	/*
		Входная точка вызывающая рекурсивную функцию на подотрезке
		от начала и до конца массива.
	*/
	ready := make(chan int, s.segmentCount)
	go s.cntSquareSegment(0, s.n, ready)
	finished := 0

	for {
		if finished == s.n {
			break
		}
		finished += <-ready
	}
}

func main() {
	var arr []int
	taskutils.ReadSlice(&arr)
	treeno := newTreeno(arr)
	treeno.CntSquares()
	taskutils.PrintSlice(treeno.arr)

}
