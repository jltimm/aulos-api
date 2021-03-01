package search

import (
	"fmt"
	"math"
	"strings"

	"../cache"
	"../common"
)

func createMap(artists []common.CacheEntry) map[string]int {
	artistsMap := make(map[string]int)
	for i := range artists {
		artistsMap[artists[i].Key] = i
	}
	return artistsMap
}

func setupMatrices(length int) ([][]int, [][]int) {
	var distanceMatrix [][]int = make([][]int, length)
	var pathMatrix [][]int = make([][]int, length)
	for i := 0; i < length; i++ {
		distanceRow := make([]int, length)
		for j := range distanceRow {
			distanceRow[j] = math.MaxInt16
		}
		distanceRow[i] = 0
		distanceMatrix[i] = distanceRow
		pathMatrix[i] = make([]int, length)
	}
	return distanceMatrix, pathMatrix
}

func FloydWarshall() {
	artists := cache.GetAllEntries()
	keyMap := createMap(artists)
	distanceMatrix, pathMatrix := setupMatrices(len(artists))
	for ind := range artists {
		artist := artists[ind]
		recommended := strings.Split(artist.Value, ",")
		for j := range recommended {
			index := keyMap[recommended[j]]
			distanceMatrix[ind][index] = 1
			distanceMatrix[index][ind] = 1
			pathMatrix[ind][index] = ind
			pathMatrix[index][ind] = index
		}
	}
	for k := 0; k < len(artists); k++ {
		for i := 0; i < len(artists); i++ {
			for j := 0; j < len(artists); j++ {
				if distanceMatrix[i][j] > (distanceMatrix[i][k] + distanceMatrix[k][j]) {
					distanceMatrix[i][j] = distanceMatrix[i][k] + distanceMatrix[k][j]
					pathMatrix[i][j] = pathMatrix[k][j]
				}
			}
		}
	}
	var path []int
	var actualPath = getPath(100, 200, pathMatrix, path)
	for anotherIndex := range actualPath {
		fmt.Println(artists[actualPath[anotherIndex]].Key)
	}
}

func getPath(i int, j int, pathMatrix [][]int, path []int) []int {
	if i == j {
		return append(path, i)
	} else {
		newPath := append(path, j)
		path = getPath(i, pathMatrix[i][j], pathMatrix, newPath)
		return path
	}
}
