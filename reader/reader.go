package reader

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Dimensões da grade do Sudoku
const (
	Rows    = 9
	Columns = 9
)

// Grade do Sudoku
type Grid [Rows][Columns]int

// Função para carregar a grade do Sudoku a partir de um arquivo
func LoadGridFromFile(filename string) (Grid, error) {
	var grid Grid

	// Lê o conteúdo do arquivo
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return grid, err
	}

	// Divide o conteúdo em linhas
	lines := strings.Split(string(content), "\n")

	// Verifica se o número de linhas é válido
	if len(lines) != Rows {
		return grid, fmt.Errorf("Número de linhas inválido")
	}

	// Preenche a grade com os números do arquivo
	for i := 0; i < Rows; i++ {
		// Divide a linha em números
		numbers := strings.Split(strings.TrimSpace(lines[i]), " ")

		// Verifica se o número de colunas é válido
		if len(numbers) != Columns {
			return grid, fmt.Errorf("Número de colunas inválido")
		}

		// Converte os números em inteiros e preenche a grade
		for j := 0; j < Columns; j++ {
			num, err := strconv.Atoi(numbers[j])
			if err != nil {
				return grid, fmt.Errorf("Número inválido na linha %d, coluna %d", i+1, j+1)
			}
			grid[i][j] = num
		}
	}

	return grid, nil
}

// Função para exibir a grade do Sudoku
func PrintGrid(grid Grid) {
	for i := 0; i < Rows; i++ {
		for j := 0; j < Columns; j++ {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Println()
	}
}
