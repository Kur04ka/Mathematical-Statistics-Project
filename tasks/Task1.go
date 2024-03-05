package tasks

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
)

func Task1() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	fmt.Println("")
	log.Println("Starting 1st task")

	// #region Наименьший и наибольший элементы выборки

	sort.Float64s(RV_X_array)
	Xmax := RV_X_array[len(RV_X_array)-1]
	Xmin := RV_X_array[0]
	fmt.Printf("Xmax: %d \tXmin: %d\n", int(Xmax), int(Xmin))

	// #endregion

	// #region Длинна интервала (Шаг)

	h := (Xmax - Xmin) / (1 + 3.322*math.Log10(float64(Nx)))
	if countDigits(int(math.Round(h))) == 1 {
		h = math.Round(h*10) / 10
		fmt.Println("Длинна интервала по правилу Стерджеса, округленная до десятых:", h)
	} else if countDigits(int(math.Round(h))) == 2 {
		h = math.Round(h)
		fmt.Println("Длинна интервала по правилу Стерджеса, округленная до целого числа:", h)
	}

	// #endregion

	// #region Создание отображения, где ключ - СВ X, а значение - Ni

	RV_X_CountMap := make(map[float64]int)
	for _, v := range RV_X_array {
		RV_X_CountMap[v]++
	}

	// #endregion

	// #region Таблица частот

	fmt.Println("Таблица частот:")
	a1 := Xmin - (h / 2)
	a2 := math.Round((a1+h)*100) / 100
	for a1 < Xmax {
		var Ni int
		for key, value := range RV_X_CountMap {
			if a1 < key && key <= a2 {
				Ni += value
			}
		}
		fmt.Println("(", a1, "; ", a2, "]\t Ni = ", Ni, "\t Wi = ", float64(Ni)/Nx,
			"\t Wi/h = ", math.Round((float64(Ni)/Nx/float64(h))*10000)/10000)

		a1 = a2
		a2 = math.Round((a1+h)*100) / 100
	}
	// #endregion

	log.Println("Ended 1st task")
}

func countDigits(number int) int {
	return len(strconv.Itoa(number))
}
