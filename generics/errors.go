package generics

type LineInsertError struct {
	value   int
	number  int
	message string
}

func (e *LineInsertError) Error() string {
	return e.message
}

type ColumnInsertError struct {
	value   int
	number  int
	message string
}

func (e *ColumnInsertError) Error() string {
	return e.message
}

type SquareInsertError struct {
	value   int
	number  int
	message string
}

func (e *SquareInsertError) Error() string {
	return e.message
}

type UnexpectedValueError struct {
	value   int
	message string
}

func (e *UnexpectedValueError) Error() string {
	return e.message
}

type InvalidPositionError struct {
	message string
}

func (e *InvalidPositionError) Error() string {
	return e.message
}

type ValueAlreadyExistsError struct {
	value   int
	message string
}

func (e *ValueAlreadyExistsError) Error() string {
	return e.message
}
