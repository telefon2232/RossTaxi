package main

import (
	"fmt"
	"os"
	"strconv"
)


const size = 23

func fillMatrix(matrix *[size][size]int) {
	matrix[0][1] = 1
	matrix[0][20] = 1
	matrix[1][0] = 1
	matrix[20][0] = 1

	matrix[1][2] = 1
	matrix[1][17] = 1
	matrix[2][1] = 1
	matrix[17][1] = 1

	matrix[2][3] = 1
	matrix[2][17] = 1
	matrix[3][2] = 1
	matrix[17][2] = 1

	matrix[3][4] = 1
	matrix[3][14] = 1
	matrix[3][16] = 1
	matrix[4][3] = 1
	matrix[14][3] = 1
	matrix[16][3] = 1

	matrix[4][5] = 1
	matrix[4][12] = 1
	matrix[5][4] = 1
	matrix[12][4] = 1
	matrix[5][6] = 1
	matrix[5][11] = 1
	matrix[6][5] = 1
	matrix[11][5] = 1

	matrix[6][7] = 1
	matrix[7][6] = 1

	matrix[7][8] = 1
	matrix[8][7] = 1

	matrix[8][9] = 1
	matrix[9][8] = 1

	matrix[9][10] = 1
	matrix[10][9] = 1

	matrix[10][11] = 1
	matrix[10][13] = 1
	matrix[11][10] = 1
	matrix[13][10] = 1

	matrix[11][12] = 1
	matrix[12][11] = 1

	matrix[12][13] = 1
	matrix[12][14] = 1
	matrix[13][12] = 1
	matrix[14][12] = 1

	matrix[13][15] = 1
	matrix[15][13] = 1

	matrix[14][15] = 1
	matrix[15][14] = 1

	matrix[15][18] = 1
	matrix[18][15] = 1

	matrix[16][17] = 1
	matrix[16][18] = 1
	matrix[17][16] = 1
	matrix[18][16] = 1

	matrix[17][19] = 1
	matrix[19][17] = 1

	matrix[18][22] = 1
	matrix[22][18] = 1

	matrix[19][20] = 1
	matrix[19][21] = 1
	matrix[20][19] = 1
	matrix[21][19] = 1

	matrix[20][21] = 1
	matrix[21][20] = 1

	matrix[21][22] = 1
	matrix[22][21] = 1
}

func Dijkstra(begin int,end int, matrix *[size][size]int) (Rows [5]int) {
	var beginIndex = begin //индекс начальной вершины
	var endIndex = end // индекс конечной вершины
	relateMatrix := [size][size]int{} //матрица связей
	minDistance := [size]int{} //минимальное расстояние
	visitedArr := [size]int{} //посещенные вершины

	var temp, minIndex, min int
	

	for i:=0; i < size; i++ {
		relateMatrix[i][i] = 0
		for j:=i+1; j < size; j++ {
			relateMatrix[i][j] = 0
			relateMatrix[j][i] = 0
		}
	}

	fillMatrix(&relateMatrix)

	//Инициализация вершин и расстояний
	for i:=0; i < size; i++ {
		minDistance[i] = 10000
		visitedArr[i] = 1
	}
	minDistance[beginIndex] = 0

	// Шаг алгоритма
	for {
		minIndex = 10000
		min = 10000
		for i:=0; i < size; i++ { // Если вершину ещё не обошли и вес меньше min
			if (visitedArr[i] == 1) && (minDistance[i] < min) { // Переприсваиваем значения
				min = minDistance[i]
				minIndex = i
			}
		}
		// Добавляем найденный минимальный вес
		// к текущему весу вершины
		// и сравниваем с текущим минимальным весом вершины
		if minIndex != 10000 {
			for i:=0; i < size; i++ {
				if relateMatrix[minIndex][i] > 0 {
					temp = min + relateMatrix[minIndex][i]
					if temp < minDistance[i] {
						minDistance[i] = temp
					}
				}
			}
			visitedArr[minIndex] = 0
		}
		if !(minIndex < 10000) {
			break
		}
	}
	// Вывод кратчайших расстояний до вершин
	//fmt.Print("Кратчайшие расстояния до вершин: ")
	//for i:=0; i < size; i++ {
	//	fmt.Println("", i)
	//	fmt.Println("", minDistance[i])
	//}

	// Восстановление пути
	visitedNodes := [size]int{} // массив посещенных вершин
	
	visitedNodes[0] = endIndex + 1 // начальный элемент - конечная вершина
	var prevNodeIndex = 1 // индекс предыдущей вершины
	var weightNode = minDistance[endIndex] // вес конечной вершины

	for endIndex != beginIndex { // пока не дошли до начальной вершины
		for i:=0; i < size; i++ { // просматриваем все вершины
			if relateMatrix[i][endIndex] != 0 { // если связь есть
				var temp = weightNode - relateMatrix[i][endIndex] // определяем вес пути из предыдущей вершины
				if temp == minDistance[i] { // если вес совпал с рассчитанным
					// значит из этой вершины и был переход
					weightNode = temp // сохраняем новый вес
					endIndex = i // сохраняем предыдущую вершину
					visitedNodes[prevNodeIndex] = i + 1 // и записываем ее в массив
					prevNodeIndex++
				}
			}
		}
	}

	Rows = [5]int{}
	for i:=prevNodeIndex - 1; i >= 0; i-- {
		Rows[i] = visitedNodes[i] - 1
		//fmt.Println("", visitedNodes[i] - 1)
	}
	return Rows
}

func main() {
	//индекс начальной вершины
	beginIndex, err := strconv.Atoi(os.Args[1]) 
	if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
    }
    // индекс конечной вершины
	endIndex, err := strconv.Atoi(os.Args[2]) 
	if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
    }
	relateMat := [size][size]int{} //матрица связей
	Row := [5]int{}
	Row = Dijkstra(beginIndex, endIndex, &relateMat)
	// Вывод пути (начальная вершина оказалась в конце массива из k элементов)
	fmt.Println("Вывод кратчайшего пути")
	for i := 0; i < 5; i++ {
		fmt.Println(" ", Row[i])
	}
}
