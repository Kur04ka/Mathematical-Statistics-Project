package tasks

import (
	"fmt"
	"log"
	"math"
	"sort"
)

func Task4() {
	fmt.Println("")
	log.Println("Starting 4th task")

	// #region

	sort.Float64s(RV_X_array)
	Xmax := RV_X_array[len(RV_X_array)-1]
	Xmin := RV_X_array[0]

	h := (Xmax - Xmin) / (1 + 3.322*math.Log10(float64(Nx)))
	if countDigits(int(math.Round(h))) == 1 {
		h = math.Round(h*10) / 10
	} else if countDigits(int(math.Round(h))) == 2 {
		h = math.Round(h)
	}

	// Отображение, где ключ - СВ X, а значение - Ni
	RV_X_CountMap := make(map[float64]int)
	for _, v := range RV_X_array {
		RV_X_CountMap[v]++
	}

	// Выборочное среднее X
	var selectiveAverageX float64
	for RV_X, Ni := range RV_X_CountMap {
		selectiveAverageX += RV_X * float64(Ni)
	}
	selectiveAverageX *= (1 / Nx)

	// Выборочная исправленная дисперсия X
	var sampleVarianceX float64
	for RV_X, Ni := range RV_X_CountMap {
		sampleVarianceX += math.Pow(RV_X, 2) * float64(Ni)
	}
	sampleVarianceX -= Nx * math.Pow(selectiveAverageX, 2)
	sampleVarianceX *= 1 / (Nx - 1)
	sampleVarianceX = math.Round(sampleVarianceX*1000) / 1000

	// Исправленное выборочное СКО X
	standardDeviationX := math.Round(math.Sqrt(sampleVarianceX)*1000) / 1000

	// #endregion

	// #region Вероятности попадания СВ X в интервал
	fmt.Println("Вероятности попадания СВ X в интервал:")
	a1 := math.Inf(-1)
	a2 := Xmin - (h / 2)
	for i := 0; float64(i) < math.Inf(1); i++ {
		var1 := (a2 - selectiveAverageX) / standardDeviationX
		var2 := (a1 - selectiveAverageX) / standardDeviationX
		fmt.Println(fmt.Sprint("p", i, " = ф(", math.Round(var1*100)/100, ") - ф(", math.Round(var2*100)/100, ")"))
		a1 = a2
		if a1 < Xmax {
			a2 = math.Round((a1+h)*100) / 100
		} else if a2 == math.Inf(1) {
			break
		} else {
			a2 = math.Inf(1)
		}
	}
	// #endregion

	log.Println("Ended 4th task")
}
