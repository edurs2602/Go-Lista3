package main

import (
	"fmt"
	"generics"
	"os"
	"reader"
)

// Função para inserir os dados passados pelo usuario ao novo grid
func insertGrid(grid reader.Grid, value int, line int, column int) (reader.Grid, error) {
	newGrid, err := generics.Insert(value, line, column, grid)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return grid, nil
	}

	return newGrid, nil
}

func main() {
	var value, line, column int

	// Verifica se foi fornecido o nome do arquivo de entrada
	if len(os.Args) < 2 {
		fmt.Println("É necessário fornecer o nome do arquivo de entrada")
		return
	}

	// Obtém o nome do arquivo de entrada
	filename := os.Args[1]

	// Carrega a grade do Sudoku a partir do arquivo
	grid, err := reader.LoadGridFromFile(filename)
	if err != nil {
		fmt.Println("Erro ao carregar a grade do Sudoku:", err)
		return
	}

	// Exibe a grade do Sudoku
	fmt.Println(" ")
	reader.PrintGrid(grid)

	// Pegando os dados fornecidos pelo usuario
	fmt.Printf("Digite o novo valor para o grid: ")
	_, err = fmt.Scanln(&value)
	if err != nil {
		return
	}

	fmt.Printf("Digite a linha: ")
	_, err = fmt.Scanln(&line)
	if err != nil {
		return
	}

	fmt.Printf("Digite a coluna: ")
	_, err = fmt.Scanln(&column)
	if err != nil {
		return
	}

	newGrid, err := insertGrid(grid, value, line, column)
	if err != nil {
		return
	}

	// Printando o novo grid
	fmt.Println("\nGrid gerado: ")
	reader.PrintGrid(newGrid)

}
