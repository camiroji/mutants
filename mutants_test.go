package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckHorizontal(t *testing.T) {

	t.Run("should find one sequence", func(t *testing.T) {
		count, ok := checkHorizontalSequence("CCCCTA")
		assert.Equal(t, false, ok)
		assert.Equal(t, 1, count)
	})

	t.Run("should find two sequences and return true", func(t *testing.T) {
		count, ok := checkHorizontalSequence("CCCCTAGGGGA")
		assert.Equal(t, true, ok)
		assert.Equal(t, 2, count)
	})

	t.Run("should not find any sequence", func(t *testing.T) {
		count, ok := checkHorizontalSequence("TCCCTAGGGCA")
		assert.Equal(t, false, ok)
		assert.Equal(t, 0, count)
	})
}

func TestCheckVertical(t *testing.T) {
	dna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

	t.Run("Should find a vertical sequence", func(t *testing.T) {
		verticalVisited := make([][]bool, 6)
		for i := 0; i < len(dna); i++ {
			verticalVisited[i] = make([]bool, 6)
		}
		ok := checkVerticalSequence(1, dna, verticalVisited, 0, 4)
		assert.Equal(t, true, ok)
		assert.Equal(t, true, verticalVisited[2][4])
		assert.Equal(t, false, verticalVisited[3][4])
	})

	t.Run("should not find a vertical sequence", func(t *testing.T) {
		verticalVisited := make([][]bool, 6)
		for i := 0; i < len(dna); i++ {
			verticalVisited[i] = make([]bool, 6)
		}
		ok := checkVerticalSequence(1, dna, verticalVisited, 1, 4)
		assert.Equal(t, false, ok)
	})

	t.Run("should not look two times in the same path", func(t *testing.T) {
		verticalVisited := make([][]bool, 6)
		for i := 0; i < len(dna); i++ {
			verticalVisited[i] = make([]bool, 6)
		}
		checkVerticalSequence(1, dna, verticalVisited, 0, 4)
		ok := checkVerticalSequence(1, dna, verticalVisited, 0, 4)
		assert.Equal(t, false, ok)
	})
}


func TestCheckDiagonal(t *testing.T) {
	dna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

	t.Run("should find a diagonal sequence", func(t *testing.T) {
		diagonalVisited := make([][]bool, 6)
		for i := 0; i < len(dna); i++ {
			diagonalVisited[i] = make([]bool, 6)
		}
		ok := checkDiagonalSequence(1, dna, diagonalVisited, 0, 0)
		assert.Equal(t, true, ok)
		assert.Equal(t, true, diagonalVisited[2][2])
		assert.Equal(t, false, diagonalVisited[3][3])
	})

	t.Run("should not found a diagonal sequence", func(t *testing.T) {
		diagonalVisited := make([][]bool, 6)
		for i := 0; i < len(dna); i++ {
			diagonalVisited[i] = make([]bool, 6)
		}
		ok := checkDiagonalSequence(1, dna, diagonalVisited, 1, 1)
		assert.Equal(t, false, ok)
	})

	t.Run("should not look two times in the same diagonal", func(t *testing.T) {
		diagonalVisited := make([][]bool, 6)
		for i := 0; i < len(dna); i++ {
			diagonalVisited[i] = make([]bool, 6)
		}
		checkDiagonalSequence(1, dna, diagonalVisited, 0, 0)
		ok := checkDiagonalSequence(1, dna, diagonalVisited, 0, 0)
		assert.Equal(t, false, ok)
	})
}

func TestIsMutant(t *testing.T) {
	t.Run("Check sequences of different letters", func(t *testing.T) {
		dna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
		ok := IsMutant(dna)
		assert.Equal(t, true, ok)
	})

	t.Run("Check not mutant", func(t *testing.T) {
		dna := []string{"ATGCGA", "CGGTGC", "TTATGT", "AGAATG", "TCCCTA", "TCACTG"}
		ok := IsMutant(dna)
		assert.Equal(t, false, ok)
	})

	t.Run("Check mutant 3 sequences with overlapping letters", func(t *testing.T) {
		dna := []string{"ATGCGA", "CGGTGC", "TTATGA", "AGAATA", "TCCCAA", "TCAAAA"}
		ok := IsMutant(dna)
		assert.Equal(t, true, ok)
	})

	t.Run("Check mutant two sequences with overlapping letters", func(t *testing.T) {
		dna := []string{"ATGCGA", "CGGTGC", "TTATGA", "AGACTA", "TCCCAA", "TCAAAA"}
		ok := IsMutant(dna)
		assert.Equal(t, true, ok)
	})
}