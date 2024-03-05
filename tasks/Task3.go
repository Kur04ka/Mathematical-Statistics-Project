package tasks

import (
	"fmt"
	"log"
	"math"
)

func Task3() {
	fmt.Println("")
	log.Println("Starting 3rd task")

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
	fmt.Println("Выборочное среднее X = ", selectiveAverageX, "\nВыборочное среднее Y = ", selectiveAverageY)

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
	fmt.Println("Выборочная исправленная дисперсия X: ", sampleVarianceX)

	// Выборочная исправленная дисперсия Y
	var sampleVarianceY float64
	for RV_Y, Ni := range RV_Y_CountMap {
		sampleVarianceY += math.Pow(RV_Y, 2) * float64(Ni)
	}
	sampleVarianceY -= Ny * math.Pow(selectiveAverageY, 2)
	sampleVarianceY *= 1 / (Ny - 1)
	sampleVarianceY = math.Round(sampleVarianceY*1000) / 1000
	fmt.Println("Выборочная исправленная дисперсия Y: ", sampleVarianceY)

	// #endregion

	// #region Исправленное выборочное СКО

	// Исправленное выборочное СКО X
	standardDeviationX := math.Round(math.Sqrt(sampleVarianceX)*1000) / 1000
	// Исправленное выборочное СКО Y
	standardDeviationY := math.Round(math.Sqrt(sampleVarianceY)*1000) / 1000
	fmt.Println("Исправленное выборочное СКО СВ X: ", standardDeviationX, "\nИсправленное выборочное СКО СВ Y: ", standardDeviationY)

	// #endregion

	log.Println("Ended 3rd task")
}

// #region TEST
// var testArray = []float64{
// 	1, 1, 2, 2, 2, 3, 3, 3, 5, 5,
// 	}

// var Ntest = float64(len(testArray))

// // Test:
// Test_CountMap := make(map[float64]int)
// for _, v := range testArray {
// 	Test_CountMap[v]++
// }

// var selectiveAverageTEST float64
// for RV_Test, Ni := range Test_CountMap {
// 	selectiveAverageTEST += RV_Test * float64(Ni)
// }
// selectiveAverageTEST *= (1 / Ntest)
// fmt.Println("Выборочное среднее: ", selectiveAverageTEST)

// var sampleVarianceTEST float64
// for RV_Test, Ni := range Test_CountMap {
// 	sampleVarianceTEST += math.Pow(RV_Test, 2) * float64(Ni)
// }
// sampleVarianceTEST -= Ntest * math.Pow(selectiveAverageTEST, 2)
// sampleVarianceTEST *= 1 / (Ntest - 1)
// fmt.Println("Выборочная исправленная дисперсия: ", sampleVarianceTEST)
// #endregion
