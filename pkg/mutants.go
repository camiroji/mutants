package main

import (
	"strings"
)

func IsMutant(dna []string) bool {
	count := 0
	var n = len(dna)
	if n < 4 {
		return false
	}
	diagonalVisited := make([][]bool, n)
	for i := 0; i < n; i++ {
		diagonalVisited[i] = make([]bool, n)
	}
	verticalVisited := make([][]bool, n)
	for i := 0; i < n; i++ {
		verticalVisited[i] = make([]bool, n)
	}

	for i, row := range dna {
		hc, ok := checkHorizontalSequence(row)
		count += hc
		if ok || count > 1 {
			return true
		}
		for j, _ := range row {
			if !diagonalVisited[i][j] {
				countDiagonal := 1
				if ok := checkDiagonalSequence(countDiagonal, dna, diagonalVisited, i, j); ok {
					count++
				}
				if count > 1 {
					return true
				}
			}
			if !verticalVisited[i][j] {
				countVertical := 1
				if ok := checkVerticalSequence(countVertical, dna, verticalVisited, i, j); ok {
					count++
				}
				if count > 1 {
					return true
				}
			}
		}
	}
	return count > 1
}

func checkHorizontalSequence(row string) (int, bool) {
	count := 0
	if count += strings.Count(row, "AAAA"); count > 1 {
		return count, true
	}
	if count += strings.Count(row, "TTTT"); count > 1 {
		return count, true
	}
	if count += strings.Count(row, "CCCC"); count > 1 {
		return count, true
	}
	count += strings.Count(row, "GGGG")
	return count, count > 1
}

func checkDiagonalSequence(count int, adn []string, visited [][]bool, i int, j int) bool {
	if count >= 4 {
		return true
	}
	visited[i][j] = true
	if i+1 < len(adn) && j+1 < len(adn) && !visited[i+1][j+1] && adn[i][j] == adn[i+1][j+1] {
		count++
		if ok := checkDiagonalSequence(count, adn, visited, i+1, j+1); ok {
			return true
		}
	}
	return false
}

func checkVerticalSequence(count int, adn []string, visited [][]bool, i int, j int) bool {
	if count >= 4 {
		return true
	}
	visited[i][j] = true
	if i+1 < len(adn) && !visited[i + 1][j] && adn[i][j] == adn[i + 1][j] {
		count++
		if ok := checkVerticalSequence(count, adn, visited, i + 1,j); ok {
			return true
		}
	}
	return false
}
