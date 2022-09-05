package main

import (
	"fmt"
	"math"
)

/*
Группирования будем добиваться путём округления до 10.
Поскольку в go взятие по модулю работает как в C
(то есть -24 % 10 == -4 а не 6) то это можно сделать путём
вычитания числа по модулю 10!
*/
func roundToGroup(x float64) float64 {
	return x - math.Mod(x, 10)
}

func main() {
	/*
		мэп срезов для хранения групп
	*/
	example := []float64{-25.0, -27.0, -21.0, 13.0, 19.0, 15.5, 24.5}
	groups := make(map[float64][]float64)
	for i := range example {
		group := roundToGroup(example[i])
		_, status := groups[group]
		if !status {
			groups[group] = make([]float64, 0)
		}
		groups[group] = append(groups[group], example[i])
	}

	/*
		красиво печатаем как в условии и без запятой в конце
	*/
	for k := range groups {
		fmt.Printf("%f: {", k)
		l := len(groups[k])
		for v := 0; v < l-1; v++ {
			fmt.Printf("%f, ", groups[k][v])
		}
		fmt.Printf("%f", groups[k][l-1])
		fmt.Println("}")
	}

}
