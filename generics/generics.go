package generics

import (
	"fmt"
	"reader"
)

func checkRow(value int, line int, table reader.Grid) bool {
	for i := 0; i < reader.Rows; i++ {
		tableValue := table[line-1][i]
		if tableValue == value {
			return true
		}
	}

	return false
}

func checkColumn(value int, column int, table reader.Grid) bool {
	for i := 0; i < reader.Columns; i++ {
		tableValue := table[i][column-1]
		if tableValue == value {
			return true
		}
	}
	return false
}

func checkSquare(value int, line int, column int, table reader.Grid) bool {
	squareSize := 3 // Tamanho de cada quadrado do Sudoku (exemplo: 3x3)

	// Coordenadas do quadrado correspondente ao elemento (line, column)
	squareRow := (line - 1) / squareSize
	squareCol := (column - 1) / squareSize

	// Verificando se há elemento igual dentro do mesmo quadrado
	for i := squareRow * squareSize; i < (squareRow+1)*squareSize; i++ {
		for j := squareCol * squareSize; j < (squareCol+1)*squareSize; j++ {
			tableValue := table[i][j]
			if tableValue == value {
				return true
			}
		}
	}

	return false
}

func Insert(value int, line int, column int, table reader.Grid) (reader.Grid, error) {
	verifier := false

	if (line > 9 || line < 1) || (column > 9 || column < 1) {
		return table, &InvalidPositionError{
			message: fmt.Sprintf("The position (%d, %d) doesn't exists in the table. Both line and column can not be above 9 or bellow 0.", line, column),
		}
	}

	//Verificando se a posição já está preenchida
	if table[line-1][column-1] != 0 {
		return table, &ValueAlreadyExistsError{
			message: fmt.Sprintf("Cannot insert '%d' in the sudoku. This position already have a value!", value),
			value:   value,
		}
	}

	//Verificando se valor a ser inserido está no invervalo de 1 a 9
	if value > 9 || value < 1 {
		return table, &UnexpectedValueError{
			message: fmt.Sprintf("Cannot insert '%d' in the sudoku. The value can not be above 9 or bellow 0.", value),
			value:   value,
		}
	}

	//Verificando se há elemento igual na linha
	verifier = checkRow(value, line, table)
	if verifier {
		return table, &LineInsertError{
			message: fmt.Sprintf("Cannot insert '%d'. This value already exists in the line.", value),
			value:   value,
		}
	}

	//Verificando se há elemento igual na coluna
	verifier = checkColumn(value, column, table)
	if verifier {
		return table, &ColumnInsertError{
			message: fmt.Sprintf("Cannot insert '%d'. This value already exists in the column.", value),
			value:   value,
		}
	}

	//Verificando se há elemento igual no mesmo quadrado
	verifier = checkSquare(value, line, column, table)
	if verifier {
		return table, &SquareInsertError{
			message: fmt.Sprintf("Cannot insert '%d'. This value already exists in the square.", value),
			value:   value,
		}
	}

	table[line-1][column-1] = value

	return table, nil
}
