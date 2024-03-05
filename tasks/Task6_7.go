package tasks

import (
	"fmt"
	"log"
	"math"
)

func Task6_7() {
	fmt.Println("")
	log.Println("Starting 6th task")

	// #region Создание отображений, где ключи - СВ X и Y, а значения - Ni

	// Отображение, где ключ - СВ X, а значение - Ni
	RV_X_CountMap := make(map[float64]int)
	for _, v := range RV_X_array {
		RV_X_CountMap[v]++
	}
	// Отображение, где ключ - СВ Y, а значение - Ni
	RV_Y_CountMap := make(map[float64]int)
	for _, v := range RV_Y_array {
		RV_Y_CountMap[v]++
	}

	// #endregion

	// #region Выборочное среднее X и Y

	// Выборочное среднее X
	var selectiveAverageX float64
	for RV_X, Ni := range RV_X_CountMap {
		selectiveAverageX += RV_X * float64(Ni)
	}
	selectiveAverageX *= (1 / Nx)

	// Выборочное среднее Y
	var selectiveAverageY float64
	for RV_Y, Ni := range RV_Y_CountMap {
		selectiveAverageY += RV_Y * float64(Ni)
	}
	selectiveAverageY *= (1 / Ny)

	// #endregion

	// #region Выборочная исправленная дисперсия X и Y

	// Выборочная исправленная дисперсия X
	var sampleVarianceX float64
	for RV_X, Ni := range RV_X_CountMap {
		sampleVarianceX += math.Pow(RV_X, 2) * float64(Ni)
	}
	sampleVarianceX -= Nx * math.Pow(selectiveAverageX, 2)
	sampleVarianceX *= 1 / (Nx - 1)
	sampleVarianceX = math.Round(sampleVarianceX*1000) / 1000

	// Выборочная исправленная дисперсия Y
	var sampleVarianceY float64
	for RV_Y, Ni := range RV_Y_CountMap {
		sampleVarianceY += math.Pow(RV_Y, 2) * float64(Ni)
	}
	sampleVarianceY -= Ny * math.Pow(selectiveAverageY, 2)
	sampleVarianceY *= 1 / (Ny - 1)
	sampleVarianceY = math.Round(sampleVarianceY*1000) / 1000

	// #endregion

	// #region Исправленное выборочное СКО

	// Исправленное выборочное СКО X
	standardDeviationX := math.Round(math.Sqrt(sampleVarianceX)*1000) / 1000
	// Исправленное выборочное СКО Y
	standardDeviationY := math.Round(math.Sqrt(sampleVarianceY)*1000) / 1000

	// #endregion

	// #region Выборочный коэффициент корреляции

	for i := 0; i < len(RV_X_array); i++ {
		correlationСoefficient += RV_X_array[i] * RV_Y_array[i]
	}
	correlationСoefficient *= 1 / Nx
	correlationСoefficient -= selectiveAverageY * selectiveAverageX
	correlationСoefficient /= standardDeviationX * standardDeviationY
	correlationСoefficient = math.Round(correlationСoefficient*1000) / 1000
	fmt.Println("Выборочный коэффициент корреляции:", correlationСoefficient)

	// #endregion

	log.Println("Ended 6th task")

	fmt.Println("")
	log.Println("Starting 7th task")

	// #region Наблюдаемое значение критерия при n = 100
	// Наблюдаемое значение критерия
	observedValue := (correlationСoefficient * math.Sqrt(Nx-2)) / math.Sqrt((1 - math.Pow(correlationСoefficient, 2)))
	observedValue = math.Round(observedValue*1000) / 1000
	fmt.Println("Наблюдаемое значение критерия при n = 100:", observedValue)
	// #endregion

	

	log.Println("Ended 7th task")
}

// Выборочный коэффициент корреляции
var correlationСoefficient float64
